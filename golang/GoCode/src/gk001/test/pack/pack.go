package main

import (
	"flag"
	"gk001/test/pack/libhello"
)


var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	libhello.Hello(name)
}
