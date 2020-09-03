package main

import (
	"sync"
)

func main() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}
	go func() {
		for {
			counter.RLock()
			_ = counter.m["some_key"]
			counter.RUnlock()
		}
	}()
	go func() {
		for {
			counter.Lock()
			counter.m["some_key"]++
			counter.Unlock()
		}
	}()
	select {}
}
