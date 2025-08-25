package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0080", "菲律宾半导体和电子工业", "https://seipi.org.ph/")

	engine.SetStartingURLs([]string{"https://seipi.org.ph/post-sitemap.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("meta[property=\"article:published_time\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Attr("content")
		ctx.Authors = append(ctx.Authors, "SEIPI - Semiconductor & Electronics Industries in the Philippines")
	})

	engine.OnHTML("div[class=\"entry-content clear\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
