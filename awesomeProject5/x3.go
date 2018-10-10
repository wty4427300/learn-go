package main

import  "fmt"

var block="shabi"
func main(){
	block :="erhuo"
	{
		block:="gun"
		fmt.Println("aaaaa",block)
	}
	fmt.Println("bbbbbbbb",block)
}