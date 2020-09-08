package main

import (
	"fmt"
)

// people包含，people0012->true,peop3l4e->true,3p4e5ople->true,epeopl0->false

func main() {
	a := "3p4e5ople"
	fmt.Println(verify(&a))
	
}


func verify(input_string *string) bool{
	var count int
	sli_in := []byte(*input_string)
	sli_peo := []byte("people")
	if len(sli_in) < len(sli_peo) {
		fmt.Println("error")
		return false
	} else {
		var num int
		l1:
		for i:=0;i<len(sli_peo);i++{
			fmt.Printf("sli_peo[i]:%c\n",sli_peo[i])
			for j:=0;j<len(sli_in);j++ {
				fmt.Printf("sli_in[j]:%c\n",sli_in[j])
				if sli_peo[i] == sli_in[j] && num <= j {
					num = j
					count++
					continue l1
				} else {
					fmt.Println("error2")
					continue
				}
			}

		}
	}
	if count == len(sli_peo) {
		fmt.Println("error3")
		return true
	}else {
		fmt.Println("error4")
		return false
	}
}
