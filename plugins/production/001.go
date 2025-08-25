package production

import (
	"fmt"
	"log"
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func FetchAndVisitArticles(engine *crawlers.WebsiteEngine, page int) {

	startRow := (page-1)*6 + 1
	endRow := page * 6

	url := fmt.Sprintf(
		"https://www.wwf.org.ph/_template/_international/articles-pages-ajax.cfm?site_id=4&lang_id=1&admin_mode=false&archive_mode=false&obj_id=7961&keywords=&month=&year=&startRow=%d&endRow=%d&contentType=article",
		startRow, endRow,
	)

	c := colly.NewCollector()

	c.OnHTML(".col-sm-6 > a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href == "" {
			return
		}

		var fullURL string
		switch {
		case strings.HasPrefix(href, "http"):
			fullURL = href
		case strings.HasPrefix(href, "/"):
			fullURL = "https://www.wwf.org.ph" + href
		default:
			fullURL = "https://www.wwf.org.ph/" + href
		}

		engine.Visit(fullURL, crawlers.News)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("请求 %s 失败: %v\n", url, err)
	})

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	c.Visit(url)
}

func init() {
	engine := crawlers.Register("001", "WWF Philippines", "https://www.wwf.org.ph/")

	engine.SetStartingURLs([]string{"https://www.wwf.org.ph/our_work/knowledge_hub/"})

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

	engine.OnLaunch(func() {
		for page := 1; page <= 100; page++ {
			FetchAndVisitArticles(engine, page)
		}
	})

	engine.OnHTML(".col-md-8.col-md-offset-2 > div, div:nth-child(2) > div > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("div:nth-child(2) > div > font > br").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
		if ctx.PageType == crawlers.News {
			ctx.Content += element.Text
		}
	})
}
