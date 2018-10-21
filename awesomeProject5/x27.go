package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	intChannels:=[3]chan int{
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
	}
	index:=rand.Intn(3)
	fmt.Println(index)
	intChannels[index]<-index
	select {
	case <-intChannels[0]:
		fmt.Println("沙皮")
	case <-intChannels[1]:
		fmt.Println("shabi")
	case <-intChannels[2]:
		fmt.Println("gundan")
	default:
		fmt.Println("没啥可接")
	}
}
