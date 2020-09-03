package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("temlp parse fail", err)
		return
	}
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    20,
	}
	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    21,
	}
	list := []string{
		"zhuqiu",
		"lanqiu",
		"yumaoqiu",
	}
	// 3.渲染模板
	//err = t.Execute(w,m1)
	err = t.Execute(w, map[string]interface{}{
		"user1": u1,
		"map1":  m1,
		"hobby": list,
	})
	if err != nil {
		fmt.Println("temlp Execute fail", err)
		return
	}
}

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数
	kua := func(name string) (string, error) {
		return name+"年轻",nil
	}

	//定义模板

	//解析模板
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	t, err := template.New("hello.tmpl").Funcs(template.FuncMap{"kua99":kua}).ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("new template failed", err)
		return
	}
	name := "小王子"
	//渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter,r *http.Request){

	t,err:=template.ParseFiles("./t.tmpl","./u1.tmpl")
	if err!=nil{
		fmt.Println("template failed",err)
		return
	}
	name := "小王子"
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo",demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server start failed,", err)
		return
	}

}
