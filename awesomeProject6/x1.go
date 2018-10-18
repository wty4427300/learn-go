package main

import "fmt"

func main(){
	s1:=make([]int,0)
	fmt.Println(cap(s1))
	for i:=1;i<=5 ;i++  {
		s1=append(s1,i)
		fmt.Println(i,len(s1),cap(s1))
	}
}
