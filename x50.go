package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	type person struct {
		id int
		name string
	}
	var p=new(person)
	p.id=99
	p.name="shapi"
	fmt.Println(unsafe.Pointer(p))
	var i=uintptr(unsafe.Pointer(p))
	var j=(*int)(unsafe.Pointer(uintptr(i)))
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(*j)
}
