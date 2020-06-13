package scrapers

import (
	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// TheVergeSearch searches for links in TheVerge
func TheVergeSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text

	properties.Message = properties.Message + "- <a href=\"" + titleURL + "\">" + title + "</a>\n\n"
}
