package main

import "fmt"

func main() {
	//m := make(map[int]int)
	var m int
	go func() {
		for {
			//read := m[1]
			read := m
			fmt.Println("read is:", read)
		}
	}()
	go func() {
		for {
			//m[2] = 2
			m = 2
		}

	}()
	fmt.Println("end!")
	select {

	}

}
