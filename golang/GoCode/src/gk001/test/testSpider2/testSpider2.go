package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func main(){
	c:= colly.NewCollector(
		colly.Async(true),
		//colly.AllowedDomains("emojipedia.org"),
		colly.AllowedDomains("hellolinux.xyz"),
		//colly.AllowURLRevisit(),//same page multiple visit
		colly.MaxDepth(1),
		)
	c.Limit(&colly.LimitRule{
		Delay: 2 * time.Second,
		RandomDelay: 1 * time.Second,
		DomainGlob: "hellolinux.xyz",
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		// Extract the link from the anchor HTML element
		link := e.Attr("href")
		fmt.Println("OnHTML",link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	//numVisited := 0
	//c.OnRequest(func(r *colly.Request){
	//	if numVisited > 10 {
	//		r.Abort()
	//	}
	//	numVisited++
	//})

	c.OnResponse(func(r *colly.Response){
		fmt.Println("OnResponse:",r.Body)
		//fmt.Println("OnResponse hello")
	})


	//c.Visit("https://emojipedia.org")
	err := c.Visit("https://hellolinux.xyz")
	dealError(err)


	c.Wait()

}


func dealError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
