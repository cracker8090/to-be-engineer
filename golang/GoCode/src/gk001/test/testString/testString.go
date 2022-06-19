package main

import (
	"fmt"
	"strings"
)

func StringPlus() string{
	var s string
	s+="昵称"+":"+"飞雪无情"+"\n"
	s+="博客"+":"+"http://www.flysnow.org/"+"\n"
	s+="微信公众号"+":"+"flysnow_org"
	return s
}

func main() {

	var builder strings.Builder
	fmt.Printf("builder len:%d,cap:%d\n", builder.Len(),builder.Cap())

	builder.WriteString("xubeiping")
	//builder.Grow(10)
	fmt.Printf("builder len:%d,cap:%d\n", builder.Len(),builder.Cap())
	fmt.Println(builder.String())

	b2:=builder
	b2.WriteString("abc")
	//strings.Reader{}


	//var s strings.Builder
	//var a strings.Reader
	//var s1 bytes.Buffer
	//var a1 bytes.Reader
	//var io io.Writer
	//
	//bytes.NewBufferString("abc")
	//s:="123456"
	//var sli []byte
	//wr,err := io.Copy(sli,s)


}
