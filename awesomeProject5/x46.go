package main

import (
	"errors"
	"fmt"
)

type operate1 func(x, y int) int

func calculate1(x int,y int,op operate1)(int,error)  {
	if op==nil{
		return 0,errors.New("没函数")
	}
	return  op(x,y),nil
}

type calculateFunc func(x int, y int) (int, error)

func ganCalculator(op operate1)  calculateFunc{
	return func(x int, y int) (int, error) {
		if op==nil{
			return 0,errors.New("没函数")
		}
		return op(x,y),nil
	}
}
func main()  {
	x,y:=1,2
	op:= func(x,y int)int {
		return x+y
	}
	result,err:=calculate1(x,y,op)
	fmt.Println(result,err)
	result,err=calculate1(x,y,nil)
	fmt.Println(result,err)
	x, y = 3, 4
	add:=ganCalculator(op)
	result, err = add(x, y)
	fmt.Println(result,err)
}