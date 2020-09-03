package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func main() {
	url := "https://www.gamersky.com/news/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err) //出现错误时直接退出
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	//NewDocumentFromReader returns a Document from an io.Reader
	html, err := goquery.NewDocumentFromReader(resp.Body)

	var newList []string
	newList = getNewsLists(html, newList)
	fmt.Println("new:", len(newList))
	newnewList := removeCopySlice(newList)
	fmt.Println("newnew:", len(newnewList))

	var wg sync.WaitGroup
	for i := 0; i < len(newnewList); i++ {
		wg.Add(1)
		go getNews(newnewList[i], &wg)
	}
	wg.Wait()

}

func removeCopySlice(list []string) []string {
	result := make([]string, 0, len(list))
	temp := map[string]struct{}{}
	for _, value := range list {
		if _, ok := temp[value]; !ok {
			temp[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

func getNewsLists(html *goquery.Document, newList []string) []string {
	html.Find("a[class=tt]").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		newList = append(newList, url)
	})
	return newList
}

type News struct {
	Title   string
	Media   string
	Url     string
	PubTime string
	Content string
}

func getNews(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		wg.Done()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error:status code %d", resp.StatusCode)
		wg.Done()
		return
	}
	html, err := goquery.NewDocumentFromReader(resp.Body)
	news := News{}

	html.Find("div[class=Mid2L_tit]>h1").Each(func(i int, selection *goquery.Selection) {
		news.Title = selection.Text()
	})
	if news.Title == "" {
		wg.Done()
		return
	}

	var tempTime string
	html.Find("div[class=detail]").Each(func(i int, selection *goquery.Selection) {
		tempTime = selection.Text()
	})
	reg := regexp.MustCompile(`\d+`)
	timeString := reg.FindAllString(tempTime, -1)
	news.PubTime = fmt.Sprintf("%s-%s-%s %s:%s:%s", timeString[0], timeString[1], timeString[2], timeString[3], timeString[4], timeString[5])

	html.Find("div[class=Mid2L_con]>p").Each(func(i int, selection *goquery.Selection) {
		news.Content = news.Content + selection.Text()
	})

	//db := mysql2.DBCon()
	//
	//stmt,err:=db.Prepare(
	//	"insert into game gamesky (`title`,`url`,`media`,`content`,`pub_time`) values (?,?,?,?,?)")
	//if err!=nil{
	//	log.Println(err)
	//	wg.Done()
	//}
	//defer stmt.Close()
	//
	//rs,err:=stmt.Exec(news.Title,news.Url,news.Media,news.Content,news.PubTime)
	//if err != nil{
	//	log.Println(err)
	//	wg.Done()
	//}
	//if id,_:=rs.LastInsertId();id>0{
	//	log.Println("插入成功")
	//}
	writeFile(news, wg)
	wg.Done()

}

//func writeFile(content string) {
func writeFile(news News, wg *sync.WaitGroup) {
	fmt.Println(news.Title)
	fmt.Println(news.Url)
	fmt.Println(news.PubTime)
	fmt.Println(time.Now().Unix())
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))
	curentPath, _ := os.Getwd()
	//dirpath := curentPath + "\\" + "game" + string(time.Now().Unix())
	dirpath := curentPath + "\\" + news.Title + strconv.FormatInt(time.Now().UnixNano(), 10)
	filePath := dirpath + ".txt"
	fmt.Println("filePath", filePath)
	//if checkFileIsExist(filePath) {
	//	err := os.Remove(filePath)
	//	log.Fatal(err)
	//}
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	n, _ := file.WriteString(news.Title)
	//var file *os.File
	//file, _ = os.Create(filePath)
	//n, _ := io.WriteString(file, news.Title)
	defer file.Close()
	if n == 0 {
		wg.Done()
		return
	}
	//io.WriteString(file, news.Url)
	//io.WriteString(file, news.PubTime)
	file.WriteString("\n")
	file.WriteString(news.Url)
	file.WriteString("\n")
	file.WriteString(news.PubTime)
	file.WriteString("\n")
	file.WriteString(news.Content)

}

func checkFileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return true
}
