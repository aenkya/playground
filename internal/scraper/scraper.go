package scraper

import (
	"context"
	"log"
	"net/http"
	"sync"

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

	for i := 0; i < s.config.MaxWorkers; i++ {
		s.wg.Add(1)
		go s.worker(ctx)
	}
	return nil
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
		depth = 0
		for _, startURL := range s.config.StartingURLs {
			if url == startURL {
				s.crawlDepth[url] = depth
				s.urlQueue <- url
				return
			}
		}
	}

	s.crawlDepth[url] = depth
	s.urlQueue <- url
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
				log.Printf("Rate limiter error: %v", err)
				continue
			}

			// process url

			// store result

			log.Printf("scraped: %s, status: %d", url, 200)
		}
	}
}
