package main

import "fmt"

func main()  {
	var cl= func(x int)(func(int)(int)) { //返回值是函数，典型案例，闭包
		var p int//与php语言不同，go可以在返回函数中修改变量的值
		return func(y int)(int){
			p++
			x++
			return x+y
		}
	}
	mycl1:=cl(1)
	m1:=mycl1(1)
	m2:=mycl1(1)
	mycl2:=cl(1)
	m3:=mycl2(1)
	fmt.Println(m1,m2,m3)
}