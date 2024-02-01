package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type Czbooks struct{}

func (s *Czbooks) Supports(url string) bool {
	return strings.HasPrefix(url, "https://czbooks.net/")
}

func (s *Czbooks) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("div.name", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("a.next-chapter", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "下一章") {
			href := e.Attr("href")
			result.MoreURL = e.Request.AbsoluteURL(href)
		}
	})

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
