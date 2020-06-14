package scrapers

import (
	"log"

	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// BloombergSearch searches for links in Bloomberg
func BloombergSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text

	properties.Message = properties.Message + "- <a href=\"" + titleURL + "\">" + title + "</a>\n\n"
	log.Println(properties.Message)
}
