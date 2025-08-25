package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0053", "University of the East Manila Campus", "www.ue.edu.ph")

	engine.SetStartingURLs([]string{
		"https://www.ue.edu.ph/mla/?s=a"})

	extractorConfig := extractors.Config{
		Author:       false,
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
	engine.OnHTML(".kingster-blog-info-wrapper>div:nth-child(1)>a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML("[rel=\"author\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".kingster-single-article-content>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("a[class=\"next page-numbers\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
