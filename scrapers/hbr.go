package scrapers

import (
	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// HBRSearch searches for links in HBR
func HBRSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text

	properties.Message = properties.Message + "- <a href=\"" + properties.Domain + titleURL + "\">" + title + "</a>\n\n"
}
