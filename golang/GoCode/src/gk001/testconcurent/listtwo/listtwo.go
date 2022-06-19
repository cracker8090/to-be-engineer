package main

import (
	"fmt"
	"sync/atomic"
)

type Person struct {
	//Money int
	Money int64
}

func main() {
	p := Person{100}
	go func() {
		//p.Money += 1000
		atomic.AddInt64(&p.Money, 1000)
	}()
	atomic.AddInt64(&p.Money, 1000)
	fmt.Printf("Money: %d\n", atomic.LoadInt64(&p.Money))
}
