package main
import ( 
 "fmt"
 "time"
) 

func main() {
 fmt.Println("Hello world!")
 duration := time.Duration(10)*time.Second
 time.Sleep(duration)
}