package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"

	"strings"
)

func init() {
	engine := crawlers.Register("ZYT0063", "National University", "https://www.national-u.edu.ph")

	engine.SetStartingURLs([]string{"https://www.national-u.edu.ph/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       false, //no author
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
		if strings.Contains(element.Text, "post") {
			engine.Visit(element.Text, crawlers.Index)
		} else if strings.Contains(element.Text, "news") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".content-pad>p,.content-pad>div>span,.content-pad>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".content-pad > span:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

}
