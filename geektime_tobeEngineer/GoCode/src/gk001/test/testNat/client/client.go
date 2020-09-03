package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var tag string

const HAND_SHAKE_MSG = "i am pass msg"

func main(){
	tag = os.Args[1]
	srcAddr := &net.UDPAddr{net.IPv4zero,9982,""}
	dstAddr := &net.UDPAddr{net.ParseIP("47.116.8.77"),9981,""}
	conn,err:=net.DialUDP("udp",srcAddr,dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	if _,err := conn.Write([]byte("hello,i'm new peer:"+tag));err!=nil{
		log.Panic(err)
	}
	data := make([]byte,1024)
	n,remoteAddr,err := conn.ReadFromUDP(data)
	if err != nil{
		fmt.Printf("error during read: %s", err)
	}
	conn.Close()
	anotherPeer := parseAddr(string(data[:n]))
	fmt.Printf("lcoal:%s server:%s another:%s\n",srcAddr,remoteAddr,anotherPeer.String())
	bidirectionHole(srcAddr,&anotherPeer)

}

func parseAddr(addr string) net.UDPAddr{
	t := strings.Split(addr,":")
	port,_ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP: net.ParseIP(t[0]),
		Port: port,
	}
}

func bidirectionHole(srcAddr *net.UDPAddr,anotherAddr *net.UDPAddr){
	conn,err := net.DialUDP("udp",srcAddr,anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	if _,err = conn.Write([]byte(HAND_SHAKE_MSG));err!=nil{
		log.Println("send handshake:",err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _,err = conn.Write([]byte("from [" + tag + "]"));err != nil{
				log.Println("send msg fail",err)
			}
		}
	}()
	for{
		data := make([]byte,1024)
		n,_,err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read:%s\n",err)
		} else {
			log.Printf("recieve：%s\n",data[:n])
		}
	}
}



