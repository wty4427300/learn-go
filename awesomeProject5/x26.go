package main

import "fmt"

type Notifier interface {
	SendInt(ch chan<- int)
}

func main()  {
	var uselessChan1=make(chan <- int,1)
	var uselessChan2=make(<- chan int,1)
	fmt.Println(uselessChan1)
	fmt.Println(uselessChan2)
	a:=getIntChan()
	for elem:=range a{
		fmt.Println(elem)
	}

}
func getIntChan()<- chan int{
	num:=5
	ch:=make(chan int ,num)
	for i:=0;i<num;i++{
		ch <- i
	}
	close(ch)
	return ch
}