package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal()
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Broadcast()
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1)
}
