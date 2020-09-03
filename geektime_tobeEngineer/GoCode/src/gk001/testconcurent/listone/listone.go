package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Money  int
	mutext sync.RWMutex
}

func (p *Person) GetMoney() int {
	p.mutext.RLock()
	defer p.mutext.RUnlock()
	money := p.Money
	return money
}

func (p *Person) AddMoney(diff int) {
	p.mutext.Lock()
	defer p.mutext.Unlock()
	p.Money += diff
}

func main() {
	p := Person{Money: 100}
	go func() {
		//p.Money += 1000
		p.AddMoney(1000)
	}()

	//fmt.Printf("Money: %d\n", p.Money)
	fmt.Printf("Money: %d\n", p.GetMoney())

}
