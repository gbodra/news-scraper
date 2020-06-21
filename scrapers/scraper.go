package scrapers

import (
	"log"

	"github.com/gbodra/news-scraper/models"
	"github.com/gocolly/colly"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Scraper genereic scraper
func Scraper(sender *tb.User, bot *tb.Bot, properties models.ScraperProperties) {
	i := 0

	c := colly.NewCollector()

	c.OnHTML(properties.SearchPath, func(e *colly.HTMLElement) {
		if i < 5 {
			switch properties.Source {
			case models.Enterprisers:
				EnterprisersSearch(e, &properties)
			case models.TechCrunch:
				TechCrunchSearch(e, &properties)
			case models.HBR:
				HBRSearch(e, &properties)
			case models.TheVerge:
				TheVergeSearch(e, &properties)
			case models.Investing:
				InvestingSearch(e, &properties)
			case models.Bloomberg:
				BloombergSearch(e, &properties)
			case models.FinancialTimes:
				FinancialTimesSearch(e, &properties)
			}
		}

		i++
	})

	c.OnScraped(func(r *colly.Response) {
		bot.Send(sender, properties.Message, &tb.SendOptions{
			DisableWebPagePreview: true,
			ParseMode:             tb.ModeHTML,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit(properties.BaseURL)
}
