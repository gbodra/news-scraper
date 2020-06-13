package scrapers

import (
	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// EnterprisersSearch searches for links in Enterprisers Project
func EnterprisersSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.ChildAttr("a", "href")
	title := e.ChildText("a")

	properties.Message = properties.Message + "- <a href=\"https://" + properties.Domain + titleURL + "\">" + title + "</a>\n\n"
}
