package fpri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("fpri", "外交政策研究中心", "https://www.fpri.org/").
		SetStartingUrls([]string{"https://www.fpri.org/sitemap_index.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".rtejustify", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})

	s.OnHTML(".author", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})
}