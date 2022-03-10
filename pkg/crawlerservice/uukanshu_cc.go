package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type UUKanShuCC struct{}

func (s *UUKanShuCC) Supports(url string) bool {
	return strings.HasPrefix(url, "https://uukanshu.cc/")
}

func (s *UUKanShuCC) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".content h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("a#linkNext", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasSuffix(href, ".html") {
			result.MoreURL = e.Request.AbsoluteURL(href)
		}
	})

	c.OnHTML("p.readcotent", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
