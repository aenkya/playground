package scraper

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/temoto/robotstxt"
	"golang.org/x/time/rate"
)

type PageResult struct {
	Error       error
	URL         string
	Title       string
	Description string
	Status      int
}

type Scraper struct {
	rateLimiter *rate.Limiter
	client      *http.Client
	config      Config
	results     []PageResult
	urlQueue    chan string
	visitedURLs map[string]bool
	crawlDepth  map[string]int

	wg           sync.WaitGroup
	resultsMutex sync.Mutex
	visitedMutex sync.Mutex

	maxDepth int
}

func NewScraper(config Config) *Scraper {
	return &Scraper{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
		rateLimiter: rate.NewLimiter(rate.Limit(config.RequestsPerSec), config.RequestBursts),
		urlQueue:    make(chan string, config.URLQueueLength),
		visitedURLs: make(map[string]bool),
		results:     []PageResult{},
		crawlDepth:  make(map[string]int),
		maxDepth:    config.MaxDepth,
	}
}

func (s *Scraper) Start(ctx context.Context) error {
	for _, url := range s.config.StartingURLs {
		s.queueURL(url)
	}

	// Create a single context with timeout for all workers
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	defer cancel()

	for i := 0; i < s.config.MaxWorkers; i++ {
		s.wg.Add(1)
		go s.worker(ctx)
	}

	s.wg.Wait()
	close(s.urlQueue) // Close the channel to signal no more URLs will be sent

	return s.saveResults()
}

func (s *Scraper) queueURL(url string) {
	s.visitedMutex.Lock()
	defer s.visitedMutex.Unlock()

	if s.visitedURLs[url] {
		return
	}

	s.visitedURLs[url] = true

	// set depth to 0 for starting URLs
	depth, exists := s.crawlDepth[url]
	if !exists {
		s.crawlDepth[url] = 1
	}

	s.crawlDepth[url] = depth
	s.urlQueue <- url
}

func getGID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))

	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		return "unknown"
	}

	return string(b[:i])
}

func (s *Scraper) worker(ctx context.Context) {
	defer s.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-s.urlQueue:
			if !ok {
				return
			}
			// wait for rate limiter
			err := s.rateLimiter.Wait(ctx)
			if err != nil {
				log.Printf("[Goroutine %s]\tRate limiter error: %v", getGID(), err)
				continue
			}

			result := s.scrapePage(ctx, url)
			s.resultsMutex.Lock()
			s.results = append(s.results, result)
			s.resultsMutex.Unlock()
		}
	}
}

func (s *Scraper) scrapePage(ctx context.Context, url string) PageResult {
	result := PageResult{URL: url}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result.Error = err
		return result
	}

	req.Header.Set("User-Agent", s.config.UserAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		result.Error = err
		return result
	}
	defer resp.Body.Close()

	result.Status = resp.StatusCode
	if resp.StatusCode != http.StatusOK {
		result.Error = fmt.Errorf("non-200 status code: %d", resp.StatusCode)
		return result
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		result.Error = err
		return result
	}

	result.Title = strings.TrimSpace((doc.Find("title")).Text())
	result.Description = strings.TrimSpace(
		doc.Find("meta[name='description']").AttrOr("content", ""),
	)

	count := 0

	doc.Find("a[href]").Each(func(_ int, sel *goquery.Selection) {
		if href, exists := sel.Attr("href"); exists {
			count++

			resolvedURL := s.resolveURL(url, href)
			if resolvedURL != "" && s.shouldCrawl(resolvedURL) {
				s.queueURL(resolvedURL)
			}
		}
	})

	return result
}

func (s *Scraper) shouldCrawl(urlStr string) bool {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Check domain restrictions
	if len(s.config.AllowedDomains) > 0 {
		domainAllowed := false

		for _, domain := range s.config.AllowedDomains {
			if strings.Contains(parsedURL.Host, domain) {
				domainAllowed = true
				break
			}
		}

		if !domainAllowed {
			return false
		}
	}

	parentDepth, exists := s.crawlDepth[urlStr]
	if !exists {
		parentURL := parsedURL.Scheme + "://" + parsedURL.Host + path.Dir(parsedURL.Path)
		if parentDepth, exists = s.crawlDepth[parentURL]; !exists {
			parentDepth = 0
		}
	}

	newDepth := parentDepth + 1
	if s.maxDepth > 0 && newDepth > s.maxDepth {
		return false
	}

	s.crawlDepth[urlStr] = newDepth

	if len(s.config.URLFilters) > 0 {
		matched := false

		for _, pattern := range s.config.URLFilters {
			if strings.Contains(urlStr, pattern) {
				matched = true
				break
			}
		}

		if !matched {
			return false
		}
	}

	for _, pattern := range s.config.ExcludeURLs {
		if strings.Contains(urlStr, pattern) {
			return false
		}
	}

	if s.config.RobotsPolicy {
		allowed, err := s.isAllowedByRobots(urlStr)
		if err != nil || !allowed {
			return false
		}
	}

	return true
}

var (
	robotsCache      = make(map[string]*robotstxt.RobotsData)
	robotsCacheMutex sync.Mutex
)

func (s *Scraper) isAllowedByRobots(urlStr string) (bool, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false, err
	}

	robotsURLStr := parsedURL.Scheme + "://" + parsedURL.Host + "/robots.txt"
	robotsURL, err := url.Parse(robotsURLStr)
	if err != nil {
		return false, fmt.Errorf("robots URL %v is not valid", robotsURLStr)
	}

	robotsCacheMutex.Lock()
	robotsData, exists := robotsCache[parsedURL.Host]
	robotsCacheMutex.Unlock()

	if !exists {
		resp, err := http.Get(robotsURL.String())
		if err != nil {
			// If we can't fetch the robots.txt, we'll assume it is allowed
			return true, nil
		}
		defer resp.Body.Close()

		robotsData, err = robotstxt.FromResponse(resp)
		if err != nil {
			return true, nil
		}

		robotsCacheMutex.Lock()
		robotsCache[parsedURL.Host] = robotsData
		robotsCacheMutex.Unlock()
	}

	agent := robotsData.FindGroup(s.config.UserAgent)

	return agent.Test(parsedURL.Path), nil
}

func (s *Scraper) resolveURL(baseURL, href string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	// Handle absolute URLs
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		parsedURL, err := url.Parse(href)
		if err != nil {
			return ""
		}

		return parsedURL.String()
	}

	// Handle relative URLs
	relativeURL, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolvedURL := base.ResolveReference(relativeURL)

	return resolvedURL.String()
}

func (s *Scraper) saveResults() error {
	file, err := os.Create(s.config.OutputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"URL", "Title", "Description", "Status", "Error"}); err != nil {
		return err
	}

	for _, result := range s.results {
		errorMsg := ""
		if result.Error != nil {
			errorMsg = result.Error.Error()
		}

		if err := writer.Write([]string{
			result.URL,
			result.Title,
			result.Description,
			fmt.Sprintf("%d", result.Status),
			errorMsg,
		}); err != nil {
			return err
		}
	}

	return nil
}
