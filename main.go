package main

import (
	"log"

	"github.com/gbodra/news-scraper/scrapers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scrapers.ScrapEnterprisers()

	scrapers.ScrapTechCrunch()

	scrapers.ScrapTheVerge()

	scrapers.ScrapHBR()
}
