package libhello

import (
	"fmt"
	"sync"
)

func Hello(name string) {
	fmt.Printf("Hello,%s!\n", name)
	var s chan string
	s <- "abc"
	fmt.Print(<-s)
	//runtime.makechan
	var l sync.Mutex
	send:=sync.NewCond(&l)
	//recv:=sync.NewCond(l.Lock())
	send.L.Lock()
}
