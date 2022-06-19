package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

//func main() {
//	runtime.GOMAXPROCS(1)
//	go func() {
//		fmt.Println("hello world")
//		// panic("hello world")  // 强制观察输出
//	}()
//	go func() {
//		for {
//			fmt.Println("aaa")  // 非内联函数，这行注释打开，将导致 hello world 的输出
//		}
//	}()
//	select {}
//}

func main() {
	//wgs := sync.WaitGroup{}
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.golang.org/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			//http.Get(url)
			fmt.Println(http.Get("http://www.baidu.com/"))
		}(url)
	}
	wg.Wait()
	//context.WithCancel()

	switch num := 8; num {
	case 1:
		fmt.Println(1)
	case 2, 4:
		fmt.Println(2, 4)
	case 8:
		fmt.Println(8)
	default:
		fmt.Println("default")
	}
	file, _ := os.Open("I:\\git_project\\to-be-engineer\\geektime_tobeEngineer\\GoCode\\src\\test1.csv")
	defer file.Close()
	r1 := csv.NewReader(file)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatal("cannot read file %+v", err)
	}
	fmt.Println(content)

}
