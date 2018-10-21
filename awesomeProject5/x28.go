package main

import (
	"fmt"
	"time"
)

func main(){
	intChan :=make(chan int,1)
	time.AfterFunc(time.Second,func(){
	close(intChan)
	})
	select {
	case _, ok:=<-intChan:
		if !ok{
			fmt.Println("meijie到")
            break
		}
	fmt.Println("接到le")
	}
}
