package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0079", "预算管理部", "https://www.dbm.gov.ph/")

	engine.SetStartingURLs([]string{"https://www.dbm.gov.ph/index.php/management-2"})

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

	engine.OnHTML("h2[itemprop=\"headline\"] > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".published > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("datetime")
		ctx.Authors = append(ctx.Authors, "Department of Budget and Management")
	})

	engine.OnHTML("div[itemprop=\"articleBody\"] span, div[itemprop=\"articleBody\"] div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("a[aria-label=\"Go to next page\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
