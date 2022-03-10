package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type Txt101 struct{}

func (s *Txt101) Supports(url string) bool {
	return strings.HasPrefix(url, "http://www.txt101.com/")
}

func (s *Txt101) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1.bname_content", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("td > a", func(e *colly.HTMLElement) {
		if e.Text == "下一页" {
			href := e.Attr("href")
			if strings.Count(href, "/") == 3 {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		result.Body = textNode(e)
	})

	return c.Visit(url)
}
