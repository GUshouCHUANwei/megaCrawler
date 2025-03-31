﻿package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2489", "AutomoBlog", "https://www.automoblog.com/")

	// engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{
		"https://www.automoblog.com/post-sitemap1.xml",
		"https://www.automoblog.com/post-sitemap5.xml",
		"https://www.automoblog.com/post-sitemap6.xml",
		"https://www.automoblog.com/post-sitemap7.xml",
		"https://www.automoblog.com/post-sitemap8.xml",
		"https://www.automoblog.com/post-sitemap9.xml",
		"https://www.automoblog.com/post-sitemap10.xml",
		"https://www.automoblog.com/post-sitemap11.xml",
		"https://www.automoblog.com/post-sitemap12.xml",
		"https://www.automoblog.com/post-sitemap13.xml",
		"https://www.automoblog.com/post-sitemap14.xml",
		"https://www.automoblog.com/post-sitemap15.xml",
		"https://www.automoblog.com/post-sitemap16.xml",
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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("figure.wp-block-image > a > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.entry-content > ul,div.entry-content > p,div.entry-content > h2, div.taxonomy-description > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
