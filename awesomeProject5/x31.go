package main

import (
	"errors"
	"fmt"
)

type a func(x,y int)int
func b(x int,y int,op a)(int ,error){
	if op==nil{
		return  0, errors.New("invalid operation")
	}
	return op(x,y),nil
}
func main()  {
	op := func(x,y int) int{
		return x+y
	}
	fmt.Println(op(1,2))
	cp, _ :=b(5,7,op)
	fmt.Println(cp)
}