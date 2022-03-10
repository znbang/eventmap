package crawlerservice

import (
	"strings"
	"testing"
)

func TestCrawlerService(t *testing.T) {
	// 先測試非最末章的 URL，再測試最末章的 URL
	urlList := [][]string{
		{"https://www.banxia.co/113_113055/25780106.html", "https://www.banxia.co/113_113055/25780225.html"},
		{"https://www.fantitxt.com/127285/read_1.html", "https://www.fantitxt.com/127285/read_122.html"},
		{"https://www.twfanti.com/DangManJiDaLaoFanCheYiHou/read_1.html", "https://www.twfanti.com/DangManJiDaLaoFanCheYiHou/read_1217.html"},
		{"https://read.aixdzs.com/339/339897/p100.html", "https://read.aixdzs.com/226/339897/p858.html"},
		{"http://big5.quanben.io/n/shengshidifei/457.html", "http://big5.quanben.io/n/shengshidifei/458.html"},
		{"http://big5.quanben5.com/n/gudaidizhupo/1560.html", "http://big5.quanben5.com/n/gudaidizhupo/1561.html"},
		{"https://www.xbiquge.cc/book/52578/36658361.html", "https://www.xbiquge.cc/book/52578/36663832.html"},
		{"https://www.ck101.org/29/29427/80930015.html", "https://www.ck101.org/29/29427/86522563.html"},
		{"https://www.imiaobige.com/read/116657/1523457.html", "https://www.imiaobige.com/read/116657/1523458.html"},
		{"http://tw.mingzw.net/mzwread/37186_9604821.html", "http://tw.mingzw.net/mzwread/37186_9606535.html"},
		{"https://www.101novel.com/ck101/147183/50696438.html", "https://www.101novel.com/ck101/147183/50696446.html"},
		{"https://www.xinqingdou.cc/89470/3230102.html", "https://www.xinqingdou.cc/89470/3230283.html"},
		{"http://www.shuquge.com/txt/30515/5704023.html", "http://www.shuquge.com/txt/30515/12136106.html"},
		{"https://www.tuhaoxs.com/book/juedaishuangjiao/7049246.html", "https://www.tuhaoxs.com/book/juedaishuangjiao/7049248.html"},
		{"http://www.txt101.com/id/12194/2851106.html", "http://www.txt101.com/id/12194/2851260.html"},
		{"https://www.wfxs.org/html/34103/11078934.html", "https://www.wfxs.org/html/34103/11078935.html"},
		{"https://www.xbiquge.cc/book/52578/36658361.html", "https://www.xbiquge.cc/book/52578/36663832.html"},
		{"https://www.xinqingdou.net/263/1415748.html", "https://www.xinqingdou.net/263/1415749.html"},
		{"http://tw.zhsxs.com/zhsread/57993_18998592.html", "http://www.zhsxs.com/zhsread/57993_19853697.html"},
		{"https://uukanshu.cc/book/16555/9937685.html", "https://uukanshu.cc/book/16555/9937686.html"},
		{"https://www.uukanshu.com/b/92/372676.html", "https://www.uukanshu.com/b/92/372677.html"},
	}

	c := New()
	for _, urls := range urlList {
		for i, url := range urls {
			var result Result

			if !c.Supports(url) {
				t.Errorf("Unsupported URL: %v", url)
				return
			}

			if err := c.Download(&result, url); err != nil {
				t.Error(err)
			}

			t.Log("MoreURL:", result.MoreURL)
			if i == 0 {
				// 非最末章的 MoreURL 不是空的
				if result.MoreURL == "" {
					t.Errorf("MoreURL is empty")
					return
				} else if !strings.HasPrefix(result.MoreURL, "http") {
					t.Errorf("Invalid MoreURL: %v", result.MoreURL)
					return
				}
			} else {
				// 最末章的 MoreURL 是空的
				if result.MoreURL != "" {
					t.Errorf("MoreURL is not empty")
					return
				}
			}

			t.Log("Title:", result.Title)
			if result.Title == "" {
				t.Errorf("Title is empty")
				return
			}

			t.Log("Body:", result.Body)
			if result.Body == "" {
				t.Errorf("Body is empty")
				return
			}
		}
	}
}
