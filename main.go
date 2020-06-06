package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domain := "enterprisersproject.com"
	message := ""

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("div.node-article > h2", func(e *colly.HTMLElement) {
		titleURL := e.ChildAttr("a", "href")
		title := e.ChildText("a")

		log.Printf("Link found: %q -> %s\n", title, domain+titleURL)

		message = message + "- <a href=\"https://" + domain + titleURL + "\">" + title + "</a>\n\n"
	})

	c.OnScraped(func(r *colly.Response) {
		queryParams := "chat_id=" + os.Getenv("CHAT_ID") + "&text=" + url.QueryEscape(message) + "&parse_mode=HTML&disable_web_page_preview=true"

		resp, err := http.Get("https://api.telegram.org/bot" + os.Getenv("TG_TOKEN") + "/sendMessage?" + queryParams)
		if err != nil {
			log.Fatalln("Error sending message to Telegram")
		}

		defer resp.Body.Close()
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit("https://enterprisersproject.com/")
}
