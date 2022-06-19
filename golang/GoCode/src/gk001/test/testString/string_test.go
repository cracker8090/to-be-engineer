package main_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func StringPlus() string{
	var s string
	s+="昵称"+":"+"hellolinux"+"\n"
	s+="博客"+":"+"http://www.hellolinux.xyz/"+"\n"
	s+="微信公众号"+":"+"hellolinux"
	return s
}

func StringFmt() string{
	return fmt.Sprint("昵称",":","hellolinux","\n","博客",":","http://www.hellolinux.xyz/","\n","微信公众号",":","hellolinux")
}

func StringJoin() string{
	s:=[]string{"昵称",":","hellolinux","\n","博客",":","http://www.hellolinux.xyz/","\n","微信公众号",":","hellolinux"}
	return strings.Join(s,"")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("hellolinux")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.hellolinux.xyz/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("hellolinux")
	return b.String()
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("hellolinux")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.hellolinux.xyz/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("hellolinux")
	return b.String()
}



func BenchmarkStringPlus(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringPlus()
	}
}

func BenchmarkStringFmt(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringFmt()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringJoin()
	}
}

func BenchmarkStringBuffer(b *testing.B){
	for i:=0;i<b.N;i++{
		StringBuffer()
	}
}

func BenchmarkStringBuilder(b *testing.B){
	for i:=0;i<b.N;i++{
		StringBuilder()
	}
	var s strings.Builder
	var a strings.Reader
	var s1 bytes.Buffer
	var a1 bytes.Reader
	var io io.Writer

	bytes.NewBufferString("abc")


}

//go test -bench=. -benchmem -benchtime=3s -cpuprofile=profile.out
//go tool pprof profile.out
//测试的拼接字符足够多
