package scrapers

import (
	"log"
	"strings"

	"github.com/gbodra/news-scraper/notification"
	"github.com/gocolly/colly"
)

// ScrapTechCrunch scraper for TechCrunch
func ScrapTechCrunch() {
	message := "<b><u>TechCrunch</u></b>\n\n"
	i := 0

	c := colly.NewCollector()

	c.OnHTML("a.post-block__title__link", func(e *colly.HTMLElement) {
		if i < 5 {
			titleURL := e.Attr("href")
			title := e.Text
			title = strings.ReplaceAll(title, "\n", "")
			title = strings.ReplaceAll(title, "\t", "")

			log.Printf("Link found: %q -> %s\n", title, titleURL)

			message = message + "- <a href=\"" + titleURL + "\">" + title + "</a>\n\n"
		}

		i++
	})

	c.OnScraped(func(r *colly.Response) {
		notification.SendMessageTelegram(message)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://techcrunch.com/")
}
