package main

import (
	"log"
	"os"
	"time"

	"github.com/gbodra/news-scraper/models"
	"github.com/gbodra/news-scraper/scrapers"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello I'm Vince Vega and I'll help you to stay updated on the latest news")
	})

	var enterprisersProperties models.ScraperProperties
	enterprisersProperties.Domain = "enterprisersproject.com"
	enterprisersProperties.Message = "<b><u>The Enterprisers Project</u></b>\n\n"
	enterprisersProperties.BaseURL = "https://enterprisersproject.com/"
	enterprisersProperties.SearchPath = "div.node-article > h2"
	enterprisersProperties.Source = models.Enterprisers

	var techCrunchProperties models.ScraperProperties
	techCrunchProperties.Message = "<b><u>TechCrunch</u></b>\n\n"
	techCrunchProperties.BaseURL = "https://techcrunch.com/"
	techCrunchProperties.SearchPath = "a.post-block__title__link"
	techCrunchProperties.Source = models.TechCrunch

	var theVergeProperties models.ScraperProperties
	theVergeProperties.Message = "<b><u>TheVerge</u></b>\n\n"
	theVergeProperties.BaseURL = "https://www.theverge.com/tech"
	theVergeProperties.SearchPath = "a[data-analytics-link=article]"
	theVergeProperties.Source = models.TheVerge

	var hbrProperties models.ScraperProperties
	hbrProperties.Domain = "https://hbr.org"
	hbrProperties.Message = "<b><u>HBR</u></b>\n\n"
	hbrProperties.BaseURL = "https://hbr.org/topic/innovation"
	hbrProperties.SearchPath = ".hed > a"
	hbrProperties.Source = models.HBR

	var investingProperties models.ScraperProperties
	investingProperties.Domain = "br.investing.com"
	investingProperties.Message = "<b><u>Investing.com</u></b>\n\n"
	investingProperties.BaseURL = "https://br.investing.com/analysis/"
	investingProperties.SearchPath = ".articleItem > div.textDiv > a"
	investingProperties.Source = models.Investing

	var bloombergProperties models.ScraperProperties
	bloombergProperties.Message = "<b><u>Bloomberg</u></b>\n\n"
	bloombergProperties.BaseURL = "https://www.bloomberg.com.br/blog/category/noticias-exclusivas/"
	bloombergProperties.SearchPath = "a[data-element=blog-post]"
	bloombergProperties.Source = models.Bloomberg

	var financialTimesProperties models.ScraperProperties
	financialTimesProperties.Domain = "www.ft.com"
	financialTimesProperties.Message = "<b><u>Financial Times</u></b>\n\n"
	financialTimesProperties.BaseURL = "https://www.ft.com/technology"
	financialTimesProperties.SearchPath = "a[data-trackable=heading-link]"
	financialTimesProperties.Source = models.FinancialTimes

	b.Handle("/enterprisers", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, enterprisersProperties)
	})

	b.Handle("/techcrunch", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, techCrunchProperties)
	})

	b.Handle("/theverge", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, theVergeProperties)
	})

	b.Handle("/hbr", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, hbrProperties)
	})

	b.Handle("/investing", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, investingProperties)
	})

	b.Handle("/bloomberg", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, bloombergProperties)
	})

	b.Handle("/ft", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, financialTimesProperties)
	})

	b.Handle("/allnews", func(m *tb.Message) {
		scrapers.Scraper(m.Sender, b, enterprisersProperties)
		scrapers.Scraper(m.Sender, b, techCrunchProperties)
		scrapers.Scraper(m.Sender, b, theVergeProperties)
		scrapers.Scraper(m.Sender, b, hbrProperties)
		scrapers.Scraper(m.Sender, b, investingProperties)
		scrapers.Scraper(m.Sender, b, bloombergProperties)
		scrapers.Scraper(m.Sender, b, financialTimesProperties)
	})

	b.Start()
}
