package main

import (
	"fmt"
	"os"
)

func main() {
	var a int = 2
	var b int = 3
	c := a + b
	fmt.Println("hello,世界 ", c)
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	fmt.Print("hello world!\n")
	fmt.Println("hello world012!")
	os.Exit(255)

}
