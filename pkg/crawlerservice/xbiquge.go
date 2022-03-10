package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type XBiQuGe struct{}

func (s *XBiQuGe) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.xbiquge.cc/")
}

func (s *XBiQuGe) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".bookname > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#link-next", func(e *colly.HTMLElement) {
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
