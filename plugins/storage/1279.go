package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1279", "停止造假", "https://www.stopfake.org/ru/glavnaya-2/")

	engine.SetStartingURLs([]string{"https://www.stopfake.org/ru/category/novosti/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".page-nav > a:last-of-type", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".td-post-featured-image img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
}
