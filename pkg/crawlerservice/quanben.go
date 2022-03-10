package crawlerservice

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

type QuanBen struct{}

func (s *QuanBen) Supports(url string) bool {
	return strings.HasPrefix(url, "http://www.quanben.io/") || strings.HasPrefix(url, "http://big5.quanben.io/")
}

func (s *QuanBen) Download(result *Result, url string) error {
	c := colly.NewCollector()
	c.DetectCharset = true
	cc := c.Clone()

	c.OnHTML(".main > h1", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".list_page a[rel=\"next\"]", func(e *colly.HTMLElement) {
		if e.Text == "下一頁" {
			href := e.Attr("href")
			if strings.HasSuffix(href, ".html") {
				result.MoreURL = e.Request.AbsoluteURL(href)
			}
		}
	})

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		e.DOM.Find("div#loading").Remove()
		result.Body = paragraphNode(e)
	})

	c.OnHTML("head > script", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "load_more") {
			return
		}

		r := regexp.MustCompile("callback=(.+)&pinyin")
		if r.MatchString(e.Text) {
			matches := r.FindStringSubmatch(e.Text)
			if len(matches) == 2 {
				e.Request.Ctx.Put("callback", matches[1])
			}
		}
	})

	c.OnHTML(".content_more .more a", func(e *colly.HTMLElement) {
		book := ""
		id := ""
		callback := e.Request.Ctx.Get("callback")

		onclick := e.Attr("onclick")
		r := regexp.MustCompile("load_more\\('(.+)','(\\d+)'\\)")
		if r.MatchString(onclick) {
			matches := r.FindStringSubmatch(onclick)
			if len(matches) == 3 {
				book = matches[1]
				id = matches[2]
			}
		}

		if callback != "" && book != "" && id != "" {
			_ = cc.Visit(e.Request.AbsoluteURL("/index.php?c=book&a=read2.jsonp&callback=" + callback + "&pinyin=" + book + "&id=" + id))
		}
	})

	cc.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", url)
	})

	cc.OnResponse(func(r *colly.Response) {
		body := string(r.Body)
		left := strings.Index(body, "(")
		right := strings.LastIndex(body, ")")
		body = body[left+1 : right]

		var model map[string]string
		if err := json.Unmarshal([]byte(body), &model); err == nil {
			body = model["content"]
			body = strings.ReplaceAll(body, "<p>", "")
			body = strings.ReplaceAll(body, "</p>", "\n")
			result.Body += body
		}
	})

	return c.Visit(url)
}
