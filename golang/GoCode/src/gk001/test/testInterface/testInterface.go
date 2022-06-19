package main

import (
	"fmt"
	"io"
	"strings"
)

//定义一个接口
type myInterface interface {
	ChangeName(string)
	SayMyName()
}

//定义一个结构体类型
type myStruct struct {
	Name string
}

//定义这个结构体的改名方法
func (m *myStruct) ChangeName(newName string) {
	m.Name = newName
}

func (m myStruct) SayMyName() {
	fmt.Println(m.Name)
}

//一个使用接口作为参数的函数
func SetName(s myInterface, name string) {
	s.ChangeName(name)
}

func main() {
	//创建这个结构体变量
	mystruct := myStruct{
		Name: "zeta",
	}

	SetName(&mystruct, "Chow")

	mystruct.SayMyName()


	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
	go func() {

	}()
	//var file os.File
	//
	//os.OpenFile()
	//os.Open()
	//net.Dial()
	//syscall.Socket()
	//var dialer net.Dialer
	//dialer.DialContext()
	//
	//net.Conn()

}
