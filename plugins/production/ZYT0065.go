package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0065", "Para", "https://www.para.org.ph")

	engine.SetStartingURLs([]string{"https://www.para.org.ph/index.html"})

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

	engine.SetTimeout(60 * time.Second)

	engine.OnHTML(".large-7>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".row>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
