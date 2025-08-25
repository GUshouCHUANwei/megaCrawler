
package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
		
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj9-5", "远东大学", "https://www.feu.edu.ph/")
	
	engine.SetStartingURLs([]string{"https://www.feu.edu.ph/university-news-and-events/"})
	
	extractorConfig := extractors.Config{
		Author:       true,//无作者
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".link-dark", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}

