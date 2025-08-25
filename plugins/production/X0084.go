package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0084", "Small Business Corporation", "https://sbcorp.gov.ph/")

	engine.SetStartingURLs([]string{"https://sbcorp.gov.ph/news/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false, //无法获取时间
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".sitebtn", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("h1[class=\"entry-title\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
		ctx.Authors = append(ctx.Authors, "SBCorp")
		ctx.PublicationTime = "None"
	})

	engine.OnHTML(".column > p, .entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".nextpostslink", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
