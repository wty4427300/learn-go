package main

import (
	"errors"
	"fmt"
)

type operate func(x,y int)int
func calculate(x int,y int,op operate)(int ,error){
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
	cp, _ :=calculate(5,7,op)
	fmt.Println(cp)
}