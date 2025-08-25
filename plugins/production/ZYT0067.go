package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0067", "unhabitat", "https://unhabitat.org.ph")

	engine.SetStartingURLs([]string{"https://unhabitat.org.ph/sitemap_index.xml"})

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

	engine.SetTimeout(60 * time.Second)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "post") && !strings.Contains(element.Text, "tag") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".elementor-widget-theme-post-content>div>p,.elementor-widget-theme-post-content > div > p> span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".elementor-post-info__item--type-date > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}
