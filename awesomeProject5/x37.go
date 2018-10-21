package main

import "fmt"


func main(){
	type funcc func(int,int)(int)  //定义函数类型
	//注意：funcc!=func(int,int)(int)

	var opt1= func(a,b int,myfunc funcc)(int) {
		return myfunc(a,b)
	}
	var add= func(a,b int)(int) { //函数是类型
		return a+b
	}
	fmt.Println(opt1(1,2,add))
	fmt.Println("====1==")

	var funptr=new(func(int,int)(int)) //定义函数指针
	//var funvar=func(int,int)(int)//报错，因为值类型必须有代码块所在的具体内存


	funptr=&add//go语言函数名不能直接作为函数地址用，只能用匿名函数的变量名取地址得到函数地址
	fmt.Println(opt1(1,2,*funptr))

}
