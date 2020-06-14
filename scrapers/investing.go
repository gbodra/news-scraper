package scrapers

import (
	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// InvestingSearch searches for links in Investing.com
func InvestingSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text

	properties.Message = properties.Message + "- <a href=\"https://" + properties.Domain + titleURL + "\">" + title + "</a>\n\n"
}
