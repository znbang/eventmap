package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type UUKanShuCom struct{}

func (s *UUKanShuCom) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.uukanshu.com/") || strings.HasPrefix(url, "https://tw.uukanshu.net/")
}

func (s *UUKanShuCom) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1#timu", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("a#next", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasSuffix(href, ".html") {
			result.MoreURL = e.Request.AbsoluteURL(href)
		}
	})

	c.OnHTML("div#contentbox", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
		if result.Body == "" {
			result.Body = textNode(e)
		}
	})

	return c.Visit(url)
}
