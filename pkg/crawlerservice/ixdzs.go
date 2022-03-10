package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type Ixdzs struct{}

func (s *Ixdzs) Supports(url string) bool {
	return strings.HasPrefix(url, "https://tw.ixdzs.com/")
}

func (s *Ixdzs) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".line > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".link > a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "下一章") {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML(".content", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
