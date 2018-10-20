package main

import "fmt"

func main(){
	var a map[int]string
	a=make(map[int]string)
	a[1]="shabi"
	a[2]="gundan"
	a[3]="nihao"
	for b,c:=range a{
		fmt.Println(b,c)
	}
	d,ok:=a[1]
	if(ok){
		fmt.Println("存在",d)
	}else{
		fmt.Println("不存在")
	}
	delete(a,2)
	e,ok1:=a[2]
	if ok1{
		fmt.Println("存在")
	}else{
		fmt.Println("不存在",e)
	}
}
