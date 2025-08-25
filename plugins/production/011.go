package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("011", " ", "https://www.actuary.org.ph/")

	engine.SetStartingURLs([]string{"https://www.actuary.org.ph/news/news-archive/"})

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

	engine.OnHTML("#news-archive-table > tbody > tr > td > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".post-entry > p > span,.post-entry > p >strong,.post-entry > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".post-entry > p > a,.post-entry > p > br").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".wp-pagenavi > a.nextpostslink", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
