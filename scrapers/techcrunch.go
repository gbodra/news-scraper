package scrapers

import (
	"strings"

	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
)

// TechCrunchSearch searches for links in TechCrunch
func TechCrunchSearch(e *colly.HTMLElement, properties *models.ScraperProperties) {
	titleURL := e.Attr("href")
	title := e.Text
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\t", "")

	properties.Message = properties.Message + "- <a href=\"" + titleURL + "\">" + title + "</a>\n\n"
}
