package crawlerservice

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

type TuHaoXS struct{}

func (s *TuHaoXS) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.tuhaoxs.com/")
}

func (s *TuHaoXS) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1.readTitle", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#header + script", func(e *colly.HTMLElement) {
		r := regexp.MustCompile(`nex = '(.+)'`)
		if r.MatchString(e.Text) {
			matches := r.FindStringSubmatch(e.Text)
			if len(matches) == 2 {
				result.MoreURL = e.Request.AbsoluteURL(matches[1])
			}
		}
	})

	c.OnHTML("#htmlContent", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
