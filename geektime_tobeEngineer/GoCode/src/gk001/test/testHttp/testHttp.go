package main

import (
	"fmt"
	"net/http"
)

func main() {
	url1 := "http://google.cn"
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The first line of response:\n%s\n", line1)

}
