package crawlerservice

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type ZhouSi struct{}

func (s *ZhouSi) Supports(url string) bool {
	return strings.HasPrefix(url, "http://www.zhsxs.com/") || strings.HasPrefix(url, "http://tw.zhsxs.com/")
}

func (s *ZhouSi) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
		r.Headers.Add("Referer", url)
	})

	c.OnHTML("td > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("div[align=\"center\"] > a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if !strings.HasSuffix(href, "_0.html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("#form1 > table > tbody > tr > td > div:nth-child(7)", func(e *colly.HTMLElement) {
		result.Body = paragraphNode(e)
	})

	return c.Visit(url)
}
