package main

import (
	"fmt"
)

func main() {
	fmt.Println("Commencing countdown.")
	//tick :=time.Tick(1*time.Second)
	//for countdown:=10;countdown>0;countdown--{
	//	fmt.Println(countdown)
	//	j<-tick
	//}
	//launch()

	//select {
	//case <-time.After(10*time.Second):
	//	Do nothing
	//case <-abort:
	//	fmt.Println("launch aborted!")
	//	return
	//}

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}

}
