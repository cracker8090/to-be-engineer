package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("hello is ", s)
	}
	fmt.Println("hello2 is ", s)
	fmt.Println(os.Args[0])
	if x := f(); x == 6 {
		fmt.Println(x)
	} else if y:= g(x);x==y{
		fmt.Println(x,y)
		fmt.Println("else")
	}

	y:="hello world"
	y+=",xu"
	//s[0] = 'x'
	fmt.Println(y)

}
func f() int {
	return 2 + 3
}
func g(x int) int {
	return x
}
