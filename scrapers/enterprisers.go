package scrapers

import (
	"log"

	"github.com/gbodra/news-scraper/notification"
	"github.com/gocolly/colly"
)

// ScrapEnterprisers scraper for the Enterprisers Project
func ScrapEnterprisers() {
	domain := "enterprisersproject.com"
	message := "<b><u>The Enterprisers Project</u></b>\n\n"
	i := 0

	c := colly.NewCollector()

	c.OnHTML("div.node-article > h2", func(e *colly.HTMLElement) {
		if i < 5 {
			titleURL := e.ChildAttr("a", "href")
			title := e.ChildText("a")

			log.Printf("Link found: %q -> %s\n", title, domain+titleURL)

			message = message + "- <a href=\"https://" + domain + titleURL + "\">" + title + "</a>\n\n"
		}

		i++
	})

	c.OnScraped(func(r *colly.Response) {
		notification.SendMessageTelegram(message)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://enterprisersproject.com/")
}
