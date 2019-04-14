package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Enter function main.")
	// 引发 panic。
	panic(errors.New("something wrong"))
	p := recover()
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function main.")
}
//错误的例子,recover并不会被执行