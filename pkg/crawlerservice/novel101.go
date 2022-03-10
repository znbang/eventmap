package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type Novel101 struct{}

func (s *Novel101) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.101novel.com/")
}

func (s *Novel101) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#thumb > a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#content > p", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
