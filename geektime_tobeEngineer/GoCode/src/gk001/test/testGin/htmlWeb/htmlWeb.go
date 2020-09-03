package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request){
	// 2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err !=nil{
		fmt.Println("temlp parse fail",err)
		return
	}
	name := "小王子"
	// 3.渲染模板
	err = t.Execute(w,name)
	if err !=nil{
		fmt.Println("temlp Execute fail",err)
		return
	}
}


func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println("http server start failed,err:%v",err)
		return
	}


}
