package interview_test

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Printf("key:%d,value:%d \n", k, *v)
	}

}
