package scrapers

import (
	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// FinancialTimesSearch searches for links in Financial Times
func FinancialTimesSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text

	properties.Message = properties.Message + "- <a href=\"" + properties.Domain + titleURL + "\">" + title + "</a>\n\n"
}
