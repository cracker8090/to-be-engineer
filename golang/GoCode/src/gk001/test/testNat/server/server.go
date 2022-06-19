package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	listener,err:=net.ListenUDP("udp",&net.UDPAddr{net.IPv4zero,9981,""})
	if err !=nil{
		fmt.Println(err)
		return
	}
	log.Printf("lcoal addr：<%s> \n",listener.LocalAddr().String())
	peers := make([]net.UDPAddr,0,2)
	data := make([]byte,1024)
	for{
		n,remoteAddr,err := listener.ReadFromUDP(data)
		if err!=nil{
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers,*remoteAddr)
		if len(peers) == 2 {
			log.Printf("UDP hole,establish %s <--> %s Connection\n", peers[0].String(), peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("transit server exit and peers can still communicate.")
			return
		}
	}
}



