package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main()  {
	var a *int16=new(int16)//取地址
	var b int16=0x0102
	a=&b
	var c  uintptr=uintptr(unsafe.Pointer(a))//unsafe.Pointer相当于void *
	var d=c+1
	var c1=(*int8)(unsafe.Pointer(c))
	var d1=(*int8)(unsafe.Pointer(d))
	fmt.Println(*c1,*d1)
	var arr=[...]int8{1,2,3,4,5,6}
	var sli=arr[2:5]
	sli=append(sli, 99)
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(sli))
	fmt.Println(sli)
	fmt.Println("===============================")
	var a1="laji"
	v,err:=strconv.Atoi(a1)
	fmt.Println(err)
	var c2=v+2
	fmt.Println(c2)
}
