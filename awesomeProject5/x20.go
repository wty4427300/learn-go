package main

import "fmt"

func main(){
	var a int =30
	var b string="a"
	switch a {
	case 10:b="a"
	case 20:b="b"
	case 30:b="c"
	}
	switch  {
	case b=="a":fmt.Println("优秀")
	case b=="b":fmt.Println("一般")
	case b=="c":fmt.Println("差蛋")
	}
}
