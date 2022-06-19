package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

// 0, 1, 2, 3, 20, 5, 6, 7,100, 9	slice
// 	  2, 3, 20,  ,  ,  ,100, 		s1
// 			4, 5, 6, 7,100,200				 	s2
