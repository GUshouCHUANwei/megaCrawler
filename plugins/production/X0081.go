package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0081", "Institute for autonomy and governance", "https://iag.org.ph/")

	engine.SetStartingURLs([]string{"https://iag.org.ph/news"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".btn.btn-default", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("span[itemprop=\"name\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".published > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	engine.OnHTML("span[style=\"font-family: arial, helvetica, sans-serif; font-size: 12pt;\"], .articleBody", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("script").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("a[title=\"Next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
