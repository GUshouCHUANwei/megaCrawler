package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("012", " ", "https://www.pbsp.org.ph/")
	
	engine.SetStartingURLs([]string{"https://www.pbsp.org.ph/news"})
	
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

	engine.OnHTML(".div-block-3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".standard-rich-text.w-richtext > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
