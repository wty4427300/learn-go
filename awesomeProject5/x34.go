package main

import "fmt"

func main(){
	var a *int8=new(int8)
	*a=10
	fmt.Println(a,*a)
}
