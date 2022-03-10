package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type Ck101 struct{}

func (s *Ck101) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.ck101.org/")
}

func (s *Ck101) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML(".info > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".page_next_preview a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if !strings.HasSuffix(href, "/0.html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
