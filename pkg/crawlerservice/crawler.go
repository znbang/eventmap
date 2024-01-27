package crawlerservice

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Result struct {
	Title   string
	Body    string
	MoreURL string
}

type Crawler interface {
	Supports(url string) bool
	Download(result *Result, url string) error
}

type CrawlerService struct {
	crawlers []Crawler
}

func New() *CrawlerService {
	return &CrawlerService{
		crawlers: []Crawler{
			&Aixdzs{},
			&BanXia{},
			&BiQuGe{},
			&Ck101{},
			&Ixdzs{},
			&MiaoBiGe{},
			&MingZW{},
			&Novel101{},
			&QingDou{},
			&QuanBen{},
			&QuanBen5{},
			&ShuQuGe{},
			&TuHaoXS{},
			&TWFanTi{},
			&Txt101{},
			&UUKanShuCC{},
			&UUKanShuCom{},
			&Wfxs{},
			&XBiQuGe{},
			&XinQingDou{},
			&ZhouSi{},
		},
	}
}

func (s *CrawlerService) Supports(url string) bool {
	for _, c := range s.crawlers {
		if c.Supports(url) {
			return true
		}
	}
	return false
}

func (s *CrawlerService) Download(result *Result, url string) error {
	for _, c := range s.crawlers {
		if c.Supports(url) {
			return c.Download(result, url)
		}
	}
	return fmt.Errorf("unsupported URL: %v", url)
}

func paragraphNode(e *colly.HTMLElement) string {
	var sb strings.Builder

	e.ForEach("p", func(i int, e *colly.HTMLElement) {
		line := strings.TrimSpace(e.Text)
		if line != "" {
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	})

	return sb.String()
}

func textNode(e *colly.HTMLElement) string {
	var sb strings.Builder

	e.DOM.Contents().Each(func(i int, s *goquery.Selection) {
		line := strings.TrimSpace(s.Text())
		if line != "" {
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	})

	return sb.String()
}
