package main

import "fmt"

type A interface {
	AAA()
}
type B interface {
	BBB()
}
type AB1 struct {
}

func (ab *AB1)AAA(){
	fmt.Println("ab1  aaa")
}
func (ab *AB1)BBB(){
	fmt.Println("ab1  bbb")
}

type AB2 struct {
}

func (ab *AB2)AAA(){
	fmt.Println("ab2  aaa")
}
func main()  {
	a:=map[string]interface{}{
		"x1":new(AB1),
		"x2":new(AB2),
	}

	for name,obj:=range a{
		f1,isA:=obj.(A) //这样可以保证强制类型转换不会宕机
		f2,isB:=obj.(B)
		fmt.Print("================")
		fmt.Println(name,"==================")
		if isA{
			f1.AAA()
		}
		if isB{
			f2.BBB()
		}

	}
}