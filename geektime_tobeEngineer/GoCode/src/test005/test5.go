package main

import (
	"errors"
	"fmt"
	"os"
)

/*func main() {
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
		<-ch
	}
	// time.Sleep(time.Second)
}*/

func main() {
	// value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	// switch 1 + 3 {
	// case value1[0], value1[1]:
	// 	fmt.Println("0 or 1")
	// case value1[2], value1[3]:
	// 	fmt.Println("2 or 3")
	// 	fallthrough
	// case value1[4], value1[5], value1[6]:
	// 	fmt.Println("4 or 5 or 6")
	// }

	// value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch value2[4] {
	// case 0, 1:
	// 	fmt.Println("0 or 1")
	// case 2, 3:
	// 	fmt.Println("2 or 3")
	// case 4, 5, 6:
	// 	fmt.Println("4 or 5 or 6")
	// }

	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
	// s := os.Open("test5.go")
	// s.close()

}
