package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly"
	"io"
	"log"
	"os"
	"regexp"
)

var fileName string

const (
	urljuejin = "https://juejin.im/post/6861785535543771149"
)

func main() {

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("something went wrong", err)
	})
	c.OnHTML(".article-title", func(e *colly.HTMLElement) {
		fileName = e.Text
	})
	c.OnHTML(".markdown-body", func(e *colly.HTMLElement) {
		reg := regexp.MustCompile(`data-`)

		html, _ := e.DOM.Html()
		//reg.ReplaceAllString(html,"")
		//fmt.Println(html)
		//writeFile(html)
		markdown := convertHTMLToMarkdown(reg.ReplaceAllString(html, ""))
		writeFile(markdown)

	})
	c.Visit(urljuejin)
	c.Wait()

}

const juejin = "juejin"

func writeFile(content string) {
	curentPath, _ := os.Getwd()
	dirpath := curentPath + "/" + juejin
	checkPath(dirpath)
	filePath := dirpath + "/" + fileName + ".md"
	fmt.Println("filePath", filePath)
	if checkFileIsExist(filePath) {
		err := os.Remove(filePath)
		log.Fatal(err)
	}
	var file *os.File
	file, _ = os.Create(filePath)
	n, _ := io.WriteString(file, content)
	file.Close()
	if n == 0 {
		return
	}
}

func checkPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		return false, nil
	}
	os.Mkdir(path, 0777)
	return false, err
}
func checkFileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return true
}

func convertHTMLToMarkdown(html string) string {
	convert := md.NewConverter("", true, nil)
	markdown, _ := convert.ConvertString(html)
	return markdown
}
