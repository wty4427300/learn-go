package main

import "fmt"

type size int
func(sz *size)dis(){
 i:=int(*sz)
 *sz=100
 fmt.Println(i)



}
func(sz size)dis1(){
	i:=int(sz)
	sz=100
	fmt.Println(i)
	fmt.Println(sz)
}
func main()  {
	var i size
	i=99
	i.dis()
	fmt.Println(i)
	i=99
	i.dis1()
	fmt.Println(i)
}
