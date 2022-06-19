package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Interger", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unknow Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
