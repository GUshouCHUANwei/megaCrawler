
package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"


)

func init() {
	engine := crawlers.Register("009", " ", "https://wmsu.edu.ph/")
	
	engine.SetStartingURLs([]string{"https://wmsu.edu.ph/?m=202508",
									"https://wmsu.edu.ph/?m=202507",
									"https://wmsu.edu.ph/?m=202506",
									"https://wmsu.edu.ph/?m=202505",
									"https://wmsu.edu.ph/?m=202504",
									"https://wmsu.edu.ph/?m=202503",
									"https://wmsu.edu.ph/?m=202502",
									"https://wmsu.edu.ph/?m=202501", 
})
	
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

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	
	engine.OnHTML(".entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".entry-content > p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("#main > nav > div > a.next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
