package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type XinQingDou struct{}

func (s *XinQingDou) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.xinqingdou.net/")
}

func (s *XinQingDou) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".bookname > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("a#A3.next", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasSuffix(href, ".html") {
			result.MoreURL = e.Request.AbsoluteURL(href)
		}
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
