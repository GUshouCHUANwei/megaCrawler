package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0062", "DUBAI", "https://www.dubaiweek.ae/")

	engine.SetStartingURLs([]string{"https://www.dubaiweek.ae/top-news/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("h3.entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("div.page-nav > a:last-of-type", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("a.tdb-author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML("div.tdb-block-inner > p > span, div.tdb-block-inner > p, div.tdb-block-inner > h4 > span, div#textareaTextHtml >p, div.tdb-block-inner > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("noscript").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
