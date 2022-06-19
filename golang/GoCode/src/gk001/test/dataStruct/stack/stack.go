package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	Items []interface{}
}

var stackMutex sync.Mutex

func (stack *Stack) New() *Stack {
	return &Stack{[]interface{}{}}
}

func (stack *Stack) PUSH(item interface{}) {
	stackMutex.Lock()
	stack.Items = append(stack.Items, item)
	stackMutex.Unlock()
}

func (stack *Stack) POP() interface{} {
	stackMutex.Lock()
	if len(stack.Items) == 0{
		return nil
	}
	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	stackMutex.Unlock()
	return item
}

func main(){
	var stack Stack
	stack.New()
	stack.PUSH(1)
	stack.PUSH(2)
	stack.PUSH(3)
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
}