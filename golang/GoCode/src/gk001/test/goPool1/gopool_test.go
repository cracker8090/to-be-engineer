package main_test

//import (
//	"fmt"
//	"testing"
//)
//
//package gopool_test

import (
	"fmt"
	"testing"
)

func TestGopool(t *testing.T){
	var i interface{}
	i =(int32(32))
	num:=i.(int32)
	fmt.Println(num)
}
