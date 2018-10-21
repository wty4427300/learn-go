package main

import "fmt"

func main(){
	ch1:=make(chan int,2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			ch1 <- 1
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()
	for{
		elem,ok:=<-ch1
		if !ok{
			fmt.Println("通道开")
		}
		fmt.Println("关了",elem)

	}
	fmt.Println("end")
}
