package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("goroutine%d 结束\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")

}
