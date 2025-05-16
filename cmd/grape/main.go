package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"enkya.org/playground/internal/scraper"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load env file: %v", err)
	}

	cfg := scraper.Config{}
	envconfig.MustProcess("grape", &cfg)
	log.Println("Starting web scraper...")

	ctx := context.Background()

	err = scraper.Start(ctx)
	if err != nil {
		log.Fatalf("Scraper error: %v", err)
	}

	log.Println("Scraping completed successfully!")
}
