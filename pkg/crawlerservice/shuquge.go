package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type ShuQuGe struct{}

func (s *ShuQuGe) Supports(url string) bool {
	return strings.HasPrefix(url, "http://www.shuquge.com/")
}

func (s *ShuQuGe) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".content > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".page_chapter a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if !strings.HasSuffix(href, "index.html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
