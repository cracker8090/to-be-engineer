package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	var mutex sync.Mutex
	ports := make([]int,0)
	for i:=80;i<50000;i++{
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			con,err:=net.DialTimeout("tcp",fmt.Sprintf("127.0.0.1:%d",port),time.Second)
			if err != nil{
				log.Printf("error:%v.port:[%d]\n",err,port)
			} else {
				con.Close()
				log.Printf("connection successful.port:[%d]\n",port)
				mutex.Lock()
				ports = append(ports,port)
				mutex.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("open port:%v",ports)

}

