package main

import "fmt"

type a123 func(interface{}) interface{}

func (a a123)call(p interface{})interface{}  {
	return a(p)
}
func incr(v interface{})interface{}  {
	return  v.(int)+1
}
func main(){
	var i a123 =incr
	v:=i.call(2)
	fmt.Println(v)
}