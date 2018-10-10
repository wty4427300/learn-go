package main

import "fmt"

func xxx(v int)(i int,j int){
	i,j=v+1,v+3
	return
}
func main()  {
	a,b:=xxx(2)
	fmt.Println(a,b)
}
