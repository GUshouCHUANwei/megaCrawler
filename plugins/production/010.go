package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("010", " ", "https://upd.edu.ph/")

	engine.SetStartingURLs([]string{"https://upd.edu.ph/category/academe/",
		"https://upd.edu.ph/category/research-highlights/",
		"https://upd.edu.ph/category/students/",
		"https://upd.edu.ph/category/upd-now/",
		"https://upd.edu.ph/category/extension/",
		"https://upd.edu.ph/category/notice/",
		"https://upd.edu.ph/category/statements/",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".recent-article > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("body > main > section:nth-child(2) > div > div.col-sm-12 > nav > ul > li:nth-child(7) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
