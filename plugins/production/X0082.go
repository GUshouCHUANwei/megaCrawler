package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0082", "Philippine Misereor Partnership Inc", "https://pmpi.org.ph/")

	engine.SetStartingURLs([]string{"https://pmpi.org.ph/statements-press-releases/"})

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

	engine.OnHTML(".elementor-widget-container > div > article .elementor-image > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".elementor-post-info__item--type-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
		ctx.Authors = append(ctx.Authors, "PMPI")
	})

	engine.OnHTML("div[dir=\"auto\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".page-numbers.next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
