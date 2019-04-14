package main

import "fmt"

func main(){
	ch1:=make(chan int,3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	elem1:=<-ch1
	elem2:=<-ch1
	fmt.Println(elem1)
	fmt.Println(elem2)
}
