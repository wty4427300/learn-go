package main

import (
	"awesomeProject1/xxx"
	"fmt"
	"unsafe"
)

type mystring struct {
	LL int16
	BB byte
	Lens int
	Pos uintptr
}
func main() {
	fmt.Println("hello word1")
    xxx.Aaa()
	var str mystring
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Alignof(str))
	fmt.Println(unsafe.Sizeof(str.Lens),unsafe.Sizeof(str.Pos),unsafe.Sizeof(str.LL))
	fmt.Println(unsafe.Offsetof(str.LL),unsafe.Offsetof(str.BB),unsafe.Offsetof(str.Lens),unsafe.Offsetof(str.Pos))
	var as [5] int32
	as[0]=1
	as[1]=2
	as[2]=3
	fmt.Println(as)
	var as2=make([]int32,2)
	as2=append(as2,3)
	as2=append(as2,4)
	fmt.Println(as2)
	var mymap=make(map[int]int32,10)
	mymap[0]=10
	fmt.Println(mymap)
	var ass=[...]int32{1,2,3,4}
	ass[1]=7
	fmt.Println(ass)
	fmt.Println("================================================")
	var a *int=new(int)
	b:=1
	a=&b
	fmt.Println(a,*a)
	fmt.Println("=====================================================")

}

