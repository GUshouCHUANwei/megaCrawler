
package production

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj8-4", "棉兰老州立大学", "https://www.msumain.edu.ph/")
	
	engine.SetStartingURLs([]string{"https://www.msumain.edu.ph/author/oipp/"})
	
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
	
	engine.OnHTML(".gdlr-core-excerpt-read-more.gdlr-core-button.gdlr-core-rectangle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".kingster-single-article-content >p:nth-of-type(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
}

