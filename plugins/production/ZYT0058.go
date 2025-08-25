package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0058", "西利曼大学", "https://su.edu.ph")

	engine.SetStartingURLs([]string{"https://su.edu.ph/category/news-update/"})

	extractorConfig := extractors.Config{
		Author:       false, //no author
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

	engine.OnHTML(".item-content>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".wp-pagenavi>a.nextpostslink", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".wpb_wrapper>p,.single-post-content-text>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".item-meta > span:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

}
