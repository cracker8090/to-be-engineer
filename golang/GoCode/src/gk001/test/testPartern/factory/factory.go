package main

import "fmt"

//最抽象的一个工厂接口
type Factory interface {
	NewTV() Televison
	NewRefrigerator() Refrigerator
}

//两个工厂都有相应的产品接口
type Televison interface {
	TVDoSomething()
}

type Refrigerator interface {
	REFDoSomething()
}

//TCL工厂
type TCLTV struct {
}

func (TCLTV) TVDoSomething() {
	fmt.Println("TCL电视在Do Something")
}

type TCLRef struct {
}

func (TCLRef) REFDoSomething() {
	fmt.Println("TCL空调在Do Something")
}

type TCLFactory struct {
}

func (TCLFactory) NewTV() Televison {
	return TCLTV{}
}

func (TCLFactory) NewRefrigerator() Refrigerator {
	return TCLRef{}
}

//美的工厂
type MediaTV struct {
}

func (MediaTV) TVDoSomething() {
	fmt.Println("美的电视在Do Something")
}

type MediaRef struct {
}

func (MediaRef) REFDoSomething() {
	fmt.Println("美的空调在Do Something")
}

type MediaFactory struct {
}

func (MediaFactory) NewTV() Televison {
	return MediaTV{}
}
func (MediaFactory) NewRefrigerator() Refrigerator {
	return MediaRef{}
}

func main() {
	var (
		factory Factory
	)
	// 这里不管是TCL工厂还是美的工厂，因为他们都实现了Factory的接口，
	// 所以这两个类都可以直接当做Factory对象来直接使用。
	factory = &TCLFactory{}
	ref := factory.NewRefrigerator()
	ref.REFDoSomething()
	tv := factory.NewTV()
	tv.TVDoSomething()

	factory = MediaFactory{}
	ref = factory.NewRefrigerator()
	ref.REFDoSomething()
	tv = factory.NewTV()
	tv.TVDoSomething()
}
