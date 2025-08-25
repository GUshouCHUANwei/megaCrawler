
package production

import (
	
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj5-2", "菲律宾海洋石油协会", "https://www.pap.ph/")
	
	engine.SetStartingURLs([]string{"https://www.pap.ph/announcements"})
	
	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".link.text-dark", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("li.uk-active+li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".col-md-7.py-md-5>div>div>img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	engine.OnHTML(".col-md-10.pt-4.text-white>font>small", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}

