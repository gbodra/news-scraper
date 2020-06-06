package scrapers

import (
	"log"

	"github.com/gbodra/news-scraper/notification"
	"github.com/gocolly/colly"
)

// ScrapHBR scraper for HBR
func ScrapHBR() {
	domain := "https://hbr.org"
	message := "<b><u>HBR</u></b>\n\n"
	i := 0

	c := colly.NewCollector()

	c.OnHTML(".hed > a", func(e *colly.HTMLElement) {
		if i < 5 {
			titleURL := e.Attr("href")
			title := e.Text

			log.Printf("Link found: %q -> %s\n", title, domain+titleURL)

			message = message + "- <a href=\"" + domain + titleURL + "\">" + title + "</a>\n\n"
		}

		i++
	})

	c.OnScraped(func(r *colly.Response) {
		notification.SendMessageTelegram(message)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://hbr.org/topic/innovation")
}
