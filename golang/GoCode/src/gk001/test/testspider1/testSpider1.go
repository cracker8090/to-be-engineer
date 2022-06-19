package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"time"
)

var visited = map[string]bool{}

const (
	hellolinux  = "hellolinux.xyz"
	hellolinux2 = "https://hellolinux.xyz/"
	baidu       = "https://www.baidu.com/"
	studygolang = "https://studygolang.com/"
)

func main() {
	//Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("hellolinux.xyz"),
		//colly.AllowedDomains("https://tding.top"),
		colly.MaxDepth(1),
	)
	//匹配模式的是该网站详情页
	detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)
	//匹配模式的是网站列表页
	listRegex, _ := regexp.Compile(`/t/\d+#\w+`)

	//所有a标签设置回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		//已访问过得详情页或列表页，跳过
		if visited[link] && (detailRegex.Match([]byte(link)) ||
			listRegex.Match([]byte(link))) {
			fmt.Println("detail or list is visited", link)
			return
		}

		//既不是列表页，页不是详情页，不是关心的内容，跳过
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			fmt.Println("detail or list not match", link)
			return
		}

		//多数网站有反爬虫策略，sleep逻辑以免封杀
		time.Sleep(2 * time.Second)
		println("match", link)

		visited[link] = true

		time.Sleep(time.Millisecond * 2)
		c.Visit(e.Request.AbsoluteURL(link))

	})

	//添加一个request回调函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit("https://hellolinux.xyz/")
	dealError(err)

	//html, err := http.Get(studygolang)
	//dealError(err)
	//fmt.Println(html)

	//time.Sleep(2 * time.Second)

}

func dealError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
