package scrapers

import (
	"log"

	"github.com/gbodra/news-scraper/notification"
	"github.com/gocolly/colly"
)

// ScrapTheVerge scraper for TheVerge
func ScrapTheVerge() {
	message := "<b><u>TheVerge</u></b>\n\n"
	i := 0

	c := colly.NewCollector()

	c.OnHTML("a[data-analytics-link=article]", func(e *colly.HTMLElement) {
		if i < 5 {
			titleURL := e.Attr("href")
			title := e.Text

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

	c.Visit("https://www.theverge.com/tech")
}
