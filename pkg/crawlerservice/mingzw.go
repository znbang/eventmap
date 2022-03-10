package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type MingZW struct{}

func (s *MingZW) Supports(url string) bool {
	return strings.HasPrefix(url, "http://tw.mingzw.net/")
}

func (s *MingZW) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("#section-free1 > div > div", func(e *colly.HTMLElement) {
		e.DOM.Find("div").Remove()
		result.Title = strings.TrimSpace(e.DOM.Text())
	})

	c.OnHTML(".title + .content > div > a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if !strings.HasSuffix(href, "_0.html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML(".title + .content", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
