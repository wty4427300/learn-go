package main

import "fmt"

func main(){
	type xxx func(int,int)int
	var xxx1= func(a int,b int,myfuncc xxx)int {
		return myfuncc(a,b)
	}
	var xxx11= func(a int,b int,myfuncc *xxx)int {
		return (*myfuncc)(a,b)
	}
	var add1= func(a,b int)int {
		return a+b
	}
	fmt.Println(xxx1(1,2,add1),"1")
	var xxx2=new(func(int,int)int)
	xxx2=&add1
	fmt.Println(xxx1(3,4,*xxx2),"2")
	var xxx3 *func(int,int)(int)=new(func(int,int)int)
	var xxx4 *xxx=new(xxx)
	xxx4=(*xxx)(&add1)
	var ff xxx=add1
	fmt.Println(xxx11(5,6,&ff),"3")
	fmt.Println()
	fmt.Println(ff(1,2))
	fmt.Println(*xxx4)
	fmt.Println(xxx3)
}
