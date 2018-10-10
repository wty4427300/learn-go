package main

import "fmt"

type MyString = string
func main(){
	var aaa MyString="shabi"//别名类型
	type MyString2 string // 注意，这里没有等号。潜在类型
	fmt.Println(aaa)
}