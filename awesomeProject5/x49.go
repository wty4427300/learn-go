package main

import "fmt"

type aaa1 struct {
	a int
	*bbb1
}
type bbb1 struct {
	b int
}
func main()  {
	var b=new(bbb1)
	var a=aaa1{1,b}
	fmt.Println(a)
}
