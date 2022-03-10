package crawlerservice

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

type TWFanTi struct{}

func (s *TWFanTi) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.twfanti.com/") || strings.HasPrefix(url, "https://www.fantitxt.com/")
}

func (s *TWFanTi) Download(result *Result, url string) error {
	re := regexp.MustCompile("本章未完，請點擊下一頁繼續閱讀！ 第\\d+頁/共\\d+頁\n")
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1 > a", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".pt-read-text", func(e *colly.HTMLElement) {
		if len(result.Body) != 0 {
			result.Body += "\n"
		}
		result.Body += re.ReplaceAllString(paragraphNode(e), "")
	})

	c.OnHTML("a.pt-nextchapter", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "下一頁") {
			href := e.Attr("href")
			if strings.Contains(href, "/read_") {
				_ = e.Request.Visit(e.Request.AbsoluteURL(href))
			}
		} else if strings.Contains(e.Text, "下一章") {
			href := e.Attr("href")
			if strings.Contains(href, "/read_") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	return c.Visit(url)
}
