package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i32 int32

func addpo() {
	fmt.Printf("addpo:%d\n", atomic.AddInt32(&i32, 5))
}

func addneg() {
	fmt.Printf("addneg:%d\n", atomic.AddInt32(&i32, -3))
}

func addVaule(delta int32) {
	for {
		//v:=i32
		v := atomic.LoadInt32(&i32)
		if atomic.CompareAndSwapInt32(&i32, v, (v + delta)) {
			break
		}
	}
}

var box6 atomic.Value

func main() {
	atomic.StoreInt32(&i32, 5)
	fmt.Println(i32)
	fmt.Println(atomic.CompareAndSwapInt32(&i32, 5, i32+1), i32)
	addVaule(4)
	fmt.Println(i32)
	//go addpo()
	//go addneg()
	slic := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		go func() {
			box6.Store(slic)
			slic[2] = 4
			fmt.Println(slic)
		}()
	}
	var s sync.Once
	s.Do(func() {
		i32 := 32
		fmt.Println(i32)
	})
	time.Sleep(time.Second * 1)

	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}
	v1 := pool.Get()
	fmt.Printf("value1:%v\n", v1)
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	v2 := pool.Get()
	fmt.Printf("value2:%v\n", v2)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("value4:%v\n", v4)
	v5 := pool.Get()
	fmt.Printf("value5:%v\n", v5)
	v6 := pool.Get()
	fmt.Printf("value5:%v\n", v6)
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	<-ctx.Done()
}
