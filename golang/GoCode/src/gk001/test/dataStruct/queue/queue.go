package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	Items []interface{}
}

var queueMutex sync.Mutex

func (queue *Queue) New() *Queue {
	return &Queue{[]interface{}{}}
}

func (queue *Queue) PUSH(item interface{}) {
	queueMutex.Lock()
	queue.Items = append(queue.Items,item)
	queueMutex.Unlock()
}

func (queue *Queue) POP() interface{} {
	queueMutex.Lock()
	if len(queue.Items) == 0{
		return nil
	}
	item := queue.Items[0]
	queue.Items = queue.Items[1:]
	queueMutex.Unlock()
	return item
}

func main(){
	var queue Queue
	queue.New()
	queue.PUSH(1)
	queue.PUSH(2)
	queue.PUSH(3)
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
}
