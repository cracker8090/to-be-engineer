package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type person struct {
	name string
	age int
}

func main(){
	//m:=make(map[int32]string)
	//m[0] = "123456"
	//m[1] = "123456"
	//m[2] = "123456"
	//m[3] = "123456"
	//m[4] = "123456"
	//m[5] = "123456"
	//for k,v:=range m{
	//	fmt.Printf("k:%v,v:%v\n",k,v)
	//}

	//p := person{
	//	"liu",
	//	20,
	//}
	//fmt.Printf("%+v\n",p)
	//fmt.Printf("%T\n",p)
	//a:=fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	//fmt.Printf("%T\n",a)
	//log.Printf("%v",m)
	//
	array :=[...]int{1,2,3,4,5,6}
	fmt.Println("%T",array)
	fmt.Println(reflect.TypeOf(array))
	bslice := array[3:]
	fmt.Println("%T",bslice)
	fmt.Println(reflect.TypeOf(bslice))
	array[0] = 11
	bslice[0]= 44
	fmt.Println(array,bslice)

	//

	//type slice struct {
	//	arr unsafe.Pointer
	//	len int
	//	cap int
	//}
	//aa := *(*)(unsafe.Pointer(&array))
	//fmt.Println(reflect.TypeOf(aa))

	//change2(bslice)
	//
	//str1 := "123456"
	//str := "123"
	//i,_:=strconv.ParseInt(str1,10,32)
	//j,_:=strconv.Atoi(str)
	//fmt.Println(i,j)
	//fmt.Println(reflect.TypeOf(i),reflect.TypeOf(j))
	//var bb int = 56
	//fmt.Println(strconv.Itoa(bb),string(bb))
	//fmt.Println(reflect.TypeOf(strconv.Itoa(bb)),reflect.TypeOf(string(bb)))

	//types.NewSlice()

	b:=[]byte("hello world")
	a:=bytes2string(b)
	b[0] = 'a'
	a = "bello world"
	fmt.Println(a,string(b))
	//listener,err:=net.ListenUDP("udp",&net.UDPAddr{net.IPv4zero,9981,nil})
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}

}
//func change1(sli []int){
//	sli[0] = 100
//}
//func change2(sli []int){
//	sli[4] = 101
//}
type StringHeader struct {
	Data uintptr
	Len  int
}
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
func bytes2string(b []byte) (string){
	//pbytes:=(*reflect.SliceHeader)(unsafe.Pointer(&b))
	//pstring:=(*reflect.StringHeader)(unsafe.Pointer(&s))
	//pstring.Data = pbytes.Data
	//pstring.Len = pbytes.Len

	sliceHead:=(*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHead:=reflect.StringHeader{
		sliceHead.Data,
		sliceHead.Len,
	}
	return *(*string)(unsafe.Pointer(&strHead))
}
func string2bytes(s string) ([]byte){
	//pbytes:=(*reflect.SliceHeader)(unsafe.Pointer(&b))
	//pstring:=(*reflect.StringHeader)(unsafe.Pointer(&s))
	//pbytes.Data = pstring.Data
	//pbytes.Len = pstring.Len
	//pbytes.Cap = pstring.Len

	stringHead:=(*reflect.StringHeader)(unsafe.Pointer(&s))
	sh:=reflect.SliceHeader{
		stringHead.Data,
		stringHead.Len,
		stringHead.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sh))
}
