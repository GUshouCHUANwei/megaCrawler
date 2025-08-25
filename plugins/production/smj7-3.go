
package production

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj7-3", "保护国际菲律宾", "https://www.conservation.org/philippines")
	
	engine.SetStartingURLs([]string{"https://www.conservation.org/sitemap/sitemap.xml"})
	
	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
        url := element.Text
        if strings.Contains(url, "blog") { // 仅采集含"blog"的链接
            engine.Visit(url, crawlers.News)
        }
    })

	engine.OnHTML(".date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}

