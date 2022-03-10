package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type BiQuGe struct{}

func (s *BiQuGe) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.biquge.com.cn/")
}

func (s *BiQuGe) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".bookname > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".bottem1 a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
