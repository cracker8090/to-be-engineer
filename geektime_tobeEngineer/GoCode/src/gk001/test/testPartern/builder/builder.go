package main

import "fmt"

//产品类
type Product struct {
	ground string
	cement string
	roof   string
}

func (p *Product) Cement() string {
	return p.cement
}
func (p *Product) SetCement(cement string) {
	p.cement = cement
}
func (p *Product) Roof() string {
	return p.roof
}
func (p *Product) SetRoof(roof string) {
	p.roof = roof
}
func (p *Product) Ground() string {
	return p.ground
}
func (p *Product) SetGround(ground string) {
	p.ground = ground
}

//抽象建造者
type Builder interface {
	BuilderGround()
	BuilderCement()
	BuilderRoof()
	BuilderProduct() *Product
}

//具体建造者
type ConcreteBuilder struct {
	p *Product
}

func (this *ConcreteBuilder) BuilderGround() {
	this.p.SetGround("建造地基")
	fmt.Println(this.p.ground)
}
func (this *ConcreteBuilder) BuilderCement() {
	this.p.SetCement("建造房子")
	fmt.Println(this.p.Cement())
}
func (this *ConcreteBuilder) BuilderRoof() {
	this.p.SetRoof("建造房顶")
	fmt.Println(this.p.Roof())
}
func (this *ConcreteBuilder) BuilderProduct() *Product {
	fmt.Println("建造完毕")
	return this.p
}

type Director struct {
	builder Builder
}

func (this *Director) SetBuilder(builder Builder) {
	this.builder = builder
}

func (this *Director) Construst() Product {
	this.builder.BuilderGround()
	this.builder.BuilderCement()
	this.builder.BuilderRoof()
	return *this.builder.BuilderProduct()
}

func main() {
	builder := &ConcreteBuilder{p: &Product{}}
	director := &Director{builder:builder}
	//director.SetBuilder(builder)
	director.Construst()

}
