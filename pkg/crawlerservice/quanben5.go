package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type QuanBen5 struct{}

func (s *QuanBen5) Supports(url string) bool {
	return strings.HasPrefix(url, "http://www.quanben5.com/") || strings.HasPrefix(url, "http://big5.quanben5.com/")
}

func (s *QuanBen5) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1.title1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#page_next > a", func(e *colly.HTMLElement) {
		if e.Text == "下一頁" {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
