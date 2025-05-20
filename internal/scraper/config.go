package scraper

import (
	"time"
)

type Config struct {
	BaseURL        string   `split_words:"true" required:"true"`
	OutputFile     string   `split_words:"true" required:"true"`
	UserAgent      string   `split_words:"true" default:"Mozilla/5.0 GoWebScraper/1.0"`
	StartingURLs   []string `required:"true"`
	AllowedDomains []string `split_words:"true" required:"true"`
	URLFilters     []string `split_words:"true"`
	ExcludeURLs    []string
	RequestsPerSec int           `split_words:"true" default:"2"`
	RequestBursts  int           `split_words:"true" default:"1"`
	MaxWorkers     int           `split_words:"true" default:"5"`
	MaxDepth       int           `split_words:"true" default:"3"`
	URLQueueLength int           `default:"1"`
	Timeout        time.Duration `default:"30s"`
	RobotsPolicy   bool          `split_words:"true" default:"true"`
}

func NewConfig() *Config {
	return &Config{}
}
