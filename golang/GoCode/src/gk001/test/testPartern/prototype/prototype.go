package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
	"sync/atomic"
)

func deepCopy(dst,src interface{}) error{
	var buf bytes.Buffer
	if err:=gob.NewEncoder(&buf).Encode(src);err!=nil{
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}


type singleton struct {

}

var instance1 *singleton

func GetInstance1() *singleton{
	if instance1 == nil{
		instance1 = &singleton{}
	}
	return instance1
}

var mu sync.Mutex

func GetInstanceMux() *singleton{
	mu.Lock()
	defer mu.Unlock()
	if instance1 == nil{
		instance1 = &singleton{}
	}
	return instance1
}


var initialized uint32

func GetInstanceAtomic() *singleton{
	if atomic.LoadUint32(&initialized) == 1 {
		return instance1
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance1 = &singleton{}
		atomic.StoreUint32(&initialized,1)
	}
	return instance1
}

var once1 sync.Once

func GetInstanceOnce() *singleton{
	// 使用Once保证创建实例的方法永远只能运行一次
	once1.Do(func() {
		instance1 = &singleton{}
	})
	return instance1
}



type ChocolateBoiler struct {
	empty  bool
	boiled bool
}
var instance *ChocolateBoiler
var once sync.Once
func GetInstance() *ChocolateBoiler {
	// 使用Once保证创建实例的方法永远只能运行一次
	once.Do(func() {
		instance = &ChocolateBoiler{true,false}
	})
	return instance
}
func (c *ChocolateBoiler) IsEmpty() bool {
	return c.empty
}
func (c *ChocolateBoiler) IsBoiled() bool {
	return c.boiled
}

func (c *ChocolateBoiler) Fill() {
	if c.empty {
		c.empty = false
		fmt.Println("容器装满了")
	}
}

func (c *ChocolateBoiler) Drain() {
	if c.empty == false && c.boiled {
		c.empty = true
		c.boiled = false
		fmt.Println("倒入模具了")
	}
}

func (c *ChocolateBoiler) Boil() {
	if c.empty == false && c.boiled == false {
		fmt.Println("巧克力煮开了")
		c.boiled = true
	}
}

func main(){
	instance = GetInstance()
	fmt.Println(instance.IsBoiled(),instance.IsEmpty())
	instance.Fill()
	instance.Boil()
	// 在倒入模具之前我们再新建一个实例，看看会不会把之前操作全都抛弃。
	instance = GetInstance()
	// 看看是否保持了煮非空的状态
	fmt.Println(instance.IsBoiled(),instance.IsEmpty())
	instance.Drain()
	fmt.Println(instance.IsBoiled(),instance.IsEmpty())
}
