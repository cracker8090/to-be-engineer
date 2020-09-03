package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//go func() {
	//	for {
	//		x, ok := <-naturals
	//		if !ok {
	//			break
	//		}
	//		squares <- x * x
	//	}
	//	close(squares)
	//}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//for {
	//	fmt.Println(<-squares)
	//}

	for x := range squares {
		fmt.Println(x)
	}
}
