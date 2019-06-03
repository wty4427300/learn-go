package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx,calcel:=context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for  {
			select {
			case <-ctx.Done():
				fmt.Print("监控停止")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(5*time.Second)
			}

		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了,通知监控停止")
	calcel()
	time.Sleep(5*time.Second)
}
