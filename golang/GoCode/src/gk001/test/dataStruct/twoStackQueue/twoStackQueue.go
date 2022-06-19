package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Stack struct {
	//Items []interface{}
	Items []int
}

var stackMutex sync.Mutex

//func (stack *Stack) New() *Stack {
//func New() *Stack {
//return &Stack{[]interface{}{}}
//}

func NewStack(col int) *Stack {
	return &Stack{Items: make([]int, 0, col)}
}

//func (stack *Stack) PUSH(item interface{}) {
func (stack *Stack) PUSH(item int) {
	stackMutex.Lock()
	stack.Items = append(stack.Items, item)
	stackMutex.Unlock()
}

//func (stack *Stack) POP() interface{} {
func (stack *Stack) POP() int {
	stackMutex.Lock()
	if len(stack.Items) == 0 {
		return 0
	}
	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	stackMutex.Unlock()
	return item
}

type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

//func (queue *Queue)New(col int) *Queue {
func New(col int) *Queue {
	return &Queue{
		stack1: NewStack(col),
		stack2: NewStack(col),
	}
}

//func (queue Queue)Push(item interface{}){
//func (queue Queue)Push(item int){
func (queue *Queue) Push(item int) {
	for len(queue.stack2.Items) > 0 {
		queue.stack1.PUSH(queue.stack2.POP())
	}
	queue.stack1.PUSH(item)
}

//func (queue *Queue)Pop() interface{}{
func (queue *Queue) Pop() int {
	for len(queue.stack1.Items) > 0 {
		queue.stack2.PUSH(queue.stack1.POP())
	}
	item := queue.stack2.POP()
	return item
}

func main() {
	//var queue Queue
	que := New(10)
	que.Push(5)
	que.Push(6)
	que.Push(7)
	fmt.Println(que.Pop())
	fmt.Println(que.Pop())

	que.Push(8)

	fmt.Println(que.Pop())
	fmt.Println(que.Pop())

	stack3 := *list.List

}
