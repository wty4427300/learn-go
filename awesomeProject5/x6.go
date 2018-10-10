package main

import (
	"fmt"
)

func main(){
	var a string
    n, a :=getTheFlag1()
    fmt.Println(a)
	fmt.Printf("%d byte(s) were written.\n", n)
}
func getTheFlag1() (i string,j string) {
	return "二货","傻逼"
}
//对ａ的重声明