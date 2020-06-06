package main

import (
	"log"
	"net/http"
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
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("div.node-article > h2", func(e *colly.HTMLElement) {
		titleURL := e.ChildAttr("a", "href")
		title := e.ChildText("a")

		log.Printf("Link found: %q -> %s\n", title, domain+titleURL)

		message := domain + titleURL
		queryParams := "chat_id=" + os.Getenv("CHAT_ID") + "&text=" + message

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
