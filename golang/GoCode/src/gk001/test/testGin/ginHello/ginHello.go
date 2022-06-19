package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request){
	//_,_ = fmt.Fprintln(w,"hello golang!")
	//_,_ = fmt.Fprintln(w,"<h1>hello golang!</h1><h1>how are you</h1>")
	b,_ := ioutil.ReadFile("./hello.txt")
	_,_ = fmt.Fprintln(w,string(b))
}

func main(){
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":8090",nil)
	if err != nil{
		fmt.Println("listen error!")
		return
	}


}




