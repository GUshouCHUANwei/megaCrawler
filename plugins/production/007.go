package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("007", " ", "https://www.xu.edu.ph/")

	engine.SetStartingURLs([]string{"https://www.xu.edu.ph/xavier-news/234-xu-ateneo-news-ay-2025-2026",
		"https://www.xu.edu.ph/xavier-news/226-xu-ateneo-news-ay-2024-2025",
		"https://www.xu.edu.ph/xavier-news/205-xu-ateneo-news-sy-2023-2024",
		"https://www.xu.edu.ph/xavier-news/194-xu-ateneo-news-sy-2020-2023",
		"https://www.xu.edu.ph/xavier-news/192-xu-ateneo-news-sy-2021-2022",
		"https://www.xu.edu.ph/xavier-news/170-xu-ateneo-news-sy-2020-2021",
		"https://www.xu.edu.ph/xavier-news/158-xu-ateneo-news-sy-2019-2020",
		"https://www.xu.edu.ph/xavier-news/140-2018-2019",
		"https://www.xu.edu.ph/xavier-news/63-2017-2018",
		"https://www.xu.edu.ph/xavier-news/55-2016-2017",
		"https://www.xu.edu.ph/xavier-news/25-2015-2016",
		"https://www.xu.edu.ph/xavier-news/24-2014-2015",
		"https://www.xu.edu.ph/xavier-news/23-2013-2014",
		"https://www.xu.edu.ph/xavier-news/22-2012-2013",
		"https://www.xu.edu.ph/xavier-news/21-2011-2012",
		"https://www.xu.edu.ph/xavier-news/20-2010-2011",
		"https://www.xu.edu.ph/xavier-news/19-2009-2010",
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

	engine.OnHTML(".list-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".item-page p+p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".item-page p+p > img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("#adminForm > div > ul > li:nth-child(12) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

}
