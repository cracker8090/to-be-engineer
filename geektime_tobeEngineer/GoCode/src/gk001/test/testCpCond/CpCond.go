package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0
	fmt.Println(runtime.GOROOT())

	// Consumer
	go func() {
		for {
			fmt.Printf("Consumer:001\n")
			cond.L.Lock()
			fmt.Printf("Consumer:002\n")
			for condition == 0 {
				fmt.Printf("Consumer:003\n")
				cond.Wait()
				fmt.Printf("Consumer:004\n")
			}
			fmt.Printf("Consumer:005\n")
			condition--
			fmt.Printf("Consumer: %d\n", condition)
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// Producer
	for {
		time.Sleep(time.Second)
		fmt.Printf("Producer:000\n")
		cond.L.Lock()
		for condition == 3 {
			fmt.Printf("Producer:001\n")
			cond.Wait()
			fmt.Printf("Producer:002\n")
		}
		condition++
		fmt.Printf("Producer: %d\n", condition)
		cond.Signal()
		cond.L.Unlock()
	}


}
