package main

import "fmt"

func main()  {
	s1:=make([]int,5,8)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)
	s2:=[]int{1,2,3,4,5,6,7,8}
	s3:=s2[3:6]
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	fmt.Println(s3)
	s4:=s3[0:cap(s3)]
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	fmt.Println(s4)
	s5:=s3[0:2]
	fmt.Println(len(s5))
	fmt.Println(cap(s5))
	fmt.Println(s5)
}
