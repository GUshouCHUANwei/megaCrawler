package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0050", "ADAMSON UNIVERSITY", "www.adamson.edu.ph")

	engine.SetStartingURLs([]string{
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2018",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2019",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2020",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2021",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2022",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2023",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2024",
		"https://www.adamson.edu.ph/v1/?page=archive&active=3&month=1&years=2025"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("h2+.dateposted", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Join(strings.Split(fulltextwithdate, "Date Posted:"), ",")
		ctx.PublicationTime = textwithdate
	})
	engine.OnHTML("h2+p~p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
