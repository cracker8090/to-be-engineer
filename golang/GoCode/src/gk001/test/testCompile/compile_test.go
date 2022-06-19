package compile_test

import (
	"fmt"
	"testing"
)

func TestCompile(t *testing.T) {
	list := new([]int)

	*list = append(*list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)

}

func Test(t *testing.T) {
	const (
		x = iota
		_
		y //2
		z = "pi"
		k        //0
		p = iota //0
		q        //1
	)
	fmt.Println(x, y, z, k, p, q)

	//var ch chan int
	//cn := make(chan int)
}

func hello(num ...int) {
	num[0] = 18
}

func TestFunc(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
	//a := 5
	//b := 8.1
	//fmt.Println(a + b)

	a := [5]int{1, 2, 3, 4, 5}
	s := a[3:4:4]
	fmt.Println(s[0])
	fmt.Println(s)

	c := [2]int{5, 6}
	d := [2]int{5, 6}
	if c == d {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	//s := make([]string,12,15)
}

func TestInterface(t *testing.T) {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func TestMap(t *testing.T) {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])

	i := -5
	j := 5
	fmt.Printf("%+d %+d", i, j)

	str := "hello"
	//str[0] = 'x'
	fmt.Println(str)

	h := 65
	fmt.Println(string(h))

}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func m(x interface{}) {

}

func TestFunction(t *testing.T) {
	fmt.Println(f1()) //1
	fmt.Println(f2()) //5
	fmt.Println(f3()) //1

}

func TestString(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	s2 = append(s2, 5, 6, 7)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s1), cap(s1))

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func TestCal(t *testing.T) {
	a := 1
	b := 2
	defer calc("A", a, calc("10", a, b))
	a = 0
	defer calc("B", a, calc("20", a, b))
	b = 1

	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}

}

type People interface {
	//Speak(string) string
	Speak(string)
}

type Student struct{}

func (stu *Student) Speak(think string) {
	fmt.Println(think)
}

func TestInherint(t *testing.T) {
	//var stu *Student
	//var pe People = new(Student)
	var peo People = &Student{}
	//var peo People = stu
	think := "speak"
	peo.Speak(think)
	//fmt.Println(peo.Speak(think))
}

func TestKV(t *testing.T) {
	myMap := map[string]int{
		"a": 1,
		"b": 1,
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

func TestRange(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
		fmt.Println(i, v)
	}
	fmt.Println("r =", r)
	fmt.Println("a =", a)
}

func TestRangeMap(t *testing.T) {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
	//var b bool
	//b=true
	//b =1
	//b=bool(1)
	//s:=make([]int,0)

	var x = []int{2: 2,3,0: 1}
	fmt.Println(x)

}



