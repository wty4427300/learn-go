package main

import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main() {
	var a int
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
	fmt.Println(unsafe.Sizeof(n),unsafe.Sizeof(a))
}
