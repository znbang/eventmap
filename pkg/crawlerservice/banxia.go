package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type BanXia struct{}

func (s *BanXia) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.banxia.co/")
}

func (s *BanXia) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("#nr_title", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".next a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasSuffix(href, ".html") {
			result.MoreURL = e.Request.AbsoluteURL(href)
		}
	})

	c.OnHTML("#nr1", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
