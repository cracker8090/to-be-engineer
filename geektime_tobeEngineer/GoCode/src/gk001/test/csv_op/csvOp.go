package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	//"github.com/axgle/mahonia"
)

func main() {
	filename := "I:\\git_project\\to-be-engineer\\geektime_tobeEngineer\\GoCode\\src\\test1.csv"
	file, err := os.Open(filename)
	//decoder := mahonia.NewDecoder("gbk")
	if err != nil {
		log.Fatalf("open file err %v", err)
	}
	defer file.Close()
	r1 := csv.NewReader(file)
	for {
		row, err := r1.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("read file error:%v", err)
		}
		if err == io.EOF {
			fmt.Println("EOF end\n")
			break
		}
		fmt.Println(row)
	}

	cntb,err:= ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil read fail:%v",err)
	}
	r2:=csv.NewReader(strings.NewReader(string((cntb))))
	fmt.Println("ioutil readfile")

	//file2, _ := os.Open(filename)
	//defer file2.Close()
	//r2 := csv.NewReader(file2)
	content, err := r2.ReadAll()
	if err != nil {
		log.Fatalf("readall err:%v", err)
	}
	for _, row := range content {
		fmt.Println(row)
	}

	newFileName := "I:\\git_project\\to-be-engineer\\geektime_tobeEngineer\\GoCode\\src\\testwrite.csv"
	filew, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		log.Fatalf("create file error:%v",err)
	}
	defer filew.Close()
	filew.Seek(0,io.SeekEnd)
	w:=csv.NewWriter(filew)
	w.UseCRLF = true
	w.Comma = ','
	writecode := []string{"1","2","3","4","5,6"}
	err=w.Write(writecode)
	if err!= nil{
		log.Fatalf("wite file fail:%v",err)
	}
	w.Flush()


	//var newContent [][]string
	//newContent = append(newContent, []string{"1", "2", "3", "4", "5", "6"})
	//newContent = append(newContent, []string{"11", "12", "13", "14", "15", "16"})
	//newContent = append(newContent, []string{"21", "22", "23", "24", "25", "26"})
	//w.WriteAll(newContent)

	bigFile,err:=os.Open(filename)
	defer bigFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf:=bufio.NewReader(bigFile)
	for{
		line,_,err:=buf.ReadLine()
		line2:= strings.TrimSpace(string(line))
		if err!= nil{
			return
		}
		fmt.Println("bigFile read,",line2)
	}

	bigFile2,err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer bigFile2.Close()
	s:=make([]byte,4096)
	for{
		switch n,err:=bigFile2.Read(s[:]);true {
		case n < 0:
			fmt.Fprintf(os.Stderr,"error read:%s\n",err.Error())
			os.Exit(1)
		case n == 0://EOF
			log.Fatalf("EOF end")
		case n > 0:
			handle(s[0:n])
		}
	}
	runtime.GOROOT()




}

//fi, err := os.Open("E:\\goTest\\CommandWindowPrint.txt")
//if err != nil {
//	return
//}
//defer fi.Close()
//decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
//fd, err := ioutil.ReadAll(decoder.NewReader(fi))
//if err != nil {
//	return
//}
//ds := strings.Split(string(fd), "\n")
//fmt.Println("ds", ds)
