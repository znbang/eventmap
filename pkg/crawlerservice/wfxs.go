package crawlerservice

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Wfxs struct{}

func (s *Wfxs) Supports(url string) bool {
	return strings.HasPrefix(url, "https://www.wfxs.org/")
}

func (s *Wfxs) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.DOM.Contents().Not("a").Text())
	})

	c.OnHTML("#thumb > a", func(e *colly.HTMLElement) {
		if e.Text == "下一章" {
			href := e.Attr("href")
			if !strings.HasSuffix(href, "/0.html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		DOM := e.DOM.Clone()
		DOM.Find("head").Remove()
		DOM.Find("title").Remove()
		DOM.Find("meta").Remove()
		DOM.Find("link").Remove()
		DOM.Find("script").Remove()
		DOM.Find("div").Remove()
		DOM.Find("ins").Remove()
		DOM.Find("h1").Remove()
		DOM.Find("br").Remove()

		var sb strings.Builder

		DOM.Find("body").Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				line := strings.TrimSpace(s.Text())
				if line != "" {
					sb.WriteString(line)
					sb.WriteString("\n")
				}
			}
		})

		result.Body = sb.String()
	})

	return c.Visit(url)
}
