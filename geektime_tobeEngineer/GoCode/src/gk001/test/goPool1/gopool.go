package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

const (
	nums = iota
	hu
	hus
)

func main() {
	defer ants.Release()
	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("free goroutines: %d\n", ants.Free())
	fmt.Printf("finish all task.\n")

	//time.Sleep(10 * time.Second)
	//fmt.Printf("running goroutines: %d\n",ants.Running())
	//fmt.Printf("free goroutines: %d\n",ants.Free())

	p,_ :=ants.NewPoolWithFunc(10,func(i interface{}){
			myFunc(i)
			wg.Done()
	})
	defer p.Release()
	for i:=0;i>runTimes;i++{
		wg.Add(1)
		_=p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running gproutines: %d\n",p.Running())
	fmt.Printf("finish all task,result is %d\n",sum)

}
