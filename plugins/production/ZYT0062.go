package production

import (
    "megaCrawler/crawlers"
    "megaCrawler/extractors"
    "github.com/gocolly/colly/v2"
)

func init() {
    engine := crawlers.Register("ZYT0062", "UPOU University", "https://www.upou.edu.ph")
    
    engine.SetStartingURLs([]string{"https://www.upou.edu.ph/news/"})
    
    extractorConfig := extractors.Config{
        Author:       false,//no author
        Image:        false,
        Language:     true,
        PublishDate:  true,
        Tags:         true,
        Text:         false,
        Title:        true,
        TextLanguage: "",
    }
    
    extractorConfig.Apply(engine)
    engine.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
        engine.Visit(element.Attr("href"), crawlers.News)
    })
    engine.OnHTML(".pagination-next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
        engine.Visit(element.Attr("href"), crawlers.Index)
    })
    engine.OnHTML(".fusion-text >p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
        ctx.Content += element.Text
    })

}

