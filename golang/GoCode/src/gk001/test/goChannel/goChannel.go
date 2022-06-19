package main

import (
	"fmt"
	"time"
)

func main() {
	//semaphore
	sema := Sema(10)
	for i := 0; i < 100; i++ {
		sema.run(getData)
	}

}

type Semaphore struct {
	c chan struct{}
}

func (s Semaphore) run(f func()) {
	s.c <- struct{}{}
	go func() {
		f()
		<-s.c
	}()
}

func Sema(len int) *Semaphore {
	return &Semaphore{c: make(chan struct{}, 10)}
}

func longRunning(message <-chan string) {
	for {
		select {
		case <-time.After(time.Minute):
			return
		case msg := <-message:
			fmt.Println(msg)
		}
	}
}

func longRunning(message <-chan string) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			return
		case msg := <-message:
			fmt.Println(msg)
			//it is important
			if !timer.Stop() {
				<-timer.C
			}
		}
	}
	//reset to reuse
	timer.Reset(time.Minute)
}

