package main

import "fmt"

func main(){
	fmt.Println(defertest1())
	fmt.Println("===============")
	defertest2()
}
func defertest1()(int){
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	return 1
}
func defertest2(){
	for i:=0;i<2;i++{
		defer fmt.Println(i)
	}

	defer fmt.Println("======")

	for i:=0;i<2;i++{
		defer func() {
			fmt.Println(i)
		}()
	}
	//分析执行顺序:
	//1.栈中压入fmt.Println(0)
	//2.栈中压入fmt.Println(1)
	//3.栈中压入fmt.Println("======")
	//4.栈中压入func() { fmt.Println(i) }()//此时i==0
	//5.栈中压入func() { fmt.Println(i) }()//此时i==1
	//6.i==2退出循环，此时i==2
	//7.函数结束，开始弹出栈
	//8.执行5，打印i值，此时i==2
	//9.执行4，打印i值，此时i==2
	//10.执行3，打印======
	//11.执行2，打印1
	//12.执行1，打印0
}
