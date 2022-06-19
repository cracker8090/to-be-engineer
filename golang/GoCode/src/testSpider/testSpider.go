package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	//创建collector
	c := colly.NewCollector(colly.AllowedDomains("juejin.im")) // 要限定域名，否则就把全网都爬下来了
	extensions.RandomUserAgent(c)                              // 使用随机的UserAgent，最好能使用代理。这样就不容易被ban
	extensions.Referer(c)                                      // 在访问的时候带上Referrer，意思就是这一次点击是从哪个页面产生的

	//事件监听，通过callback执行事件处理
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		e.Request.Visit(e.Attr("href"))
	})
	//启动网页访问
	c.Visit("https://juejin.im/")

}
