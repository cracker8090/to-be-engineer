package main

import (
	"fmt"
	//"log"
	"time"
	//"github.com/robfig/cron"
)

// people包含，people0012->true,peop3l4e->true,3p4e5ople->true,epeopl0->false

func main() {
	t1 := time.Now()
	a := "3p4e5ople"
	fmt.Println(verify(&a))

	tt := time.Now()
	fmt.Println(tt.Sub(t1))

	//c := cron.New()
	//
	//c.AddFunc("0 */1 * * * *", func() {
	//	log.Println("every 1 seconds executing")
	//	log.Println("add func")
	//})
	//
	//go c.Start()
	//defer c.Stop()

	//i := 0
	//c1 := cron.New()
	//spec := "*/5 * * * * ?"
	//
	//c1.AddFunc(spec, func() {
	//	i++
	//	log.Println("cron running:", i)
	//})
	//
	//c1.AddJob(spec, TestJob{})
	//c1.AddJob(spec, Test2Job{})

	//select {
	//case <-time.After(time.Second * 10):
	//	return
	//}
	//select {
	//}

	//go tickerTest1()
	//tickerTest2()

	ticker1 := TickerTest()
	time.Sleep(20 * time.Second)
	ticker1.Stop()
}

func TickerTest() *time.Ticker {
	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		defer func() {
			fmt.Println("defer Ticker1 stop")
		}()
		for range ticker.C {
			fmt.Println("Ticker1······")
		}
		fmt.Println("Ticker1 stop")
	}(ticker)

	return ticker
}

func tickerTest1() {
	ticker := *time.NewTicker(time.Second)
	count := 0
	go func() {
		time.Sleep(3 * time.Second)
		ticker.Stop()
	}()
	for range ticker.C {
		count++
		fmt.Println("tickerTest1：", count)
	}
}

func tickerTest2() {
	ticker := time.NewTicker(time.Second)
	count := 0
	go func() {
		time.Sleep(3 * time.Second)
		ticker.Stop()
	}()
	for range ticker.C {
		count++
		fmt.Println("tickerTest2：", count)
	}
}

func verify(input_string *string) bool {
	var count int
	sli_in := []byte(*input_string)
	sli_peo := []byte("people")
	if len(sli_in) < len(sli_peo) {
		fmt.Println("error")
		return false
	} else {
		var num int
	l1:
		for i := 0; i < len(sli_peo); i++ {
			fmt.Printf("sli_peo[i]:%c\n", sli_peo[i])
			for j := 0; j < len(sli_in); j++ {
				fmt.Printf("sli_in[j]:%c\n", sli_in[j])
				if sli_peo[i] == sli_in[j] && num <= j {
					num = j
					count++
					continue l1
				} else {
					fmt.Println("error2")
					continue
				}
			}

		}
	}
	if count == len(sli_peo) {
		fmt.Println("error3")
		return true
	} else {
		fmt.Println("error4")
		return false
	}
}

type TestJob struct {
}

func (t TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (t Test2Job) Run() {
	fmt.Println("testJob2...")
}
