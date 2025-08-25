
package production

import (
	"time"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
			
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj3-1", "菲律宾游艇俱乐部", "https://philippineyachting.com/")
	
	engine.SetStartingURLs([]string{"https://philippineyachting.com/insights/"})
	
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

	engine.OnHTML(".elementor-post__title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
		time.Sleep(2* time.Second) 
	})

}

