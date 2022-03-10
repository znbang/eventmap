package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type MiaoBiGe struct{}

func (s *MiaoBiGe) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.imiaobige.com/")
}

func (s *MiaoBiGe) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".title > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#content + .jump > a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
