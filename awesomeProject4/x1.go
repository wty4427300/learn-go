package main

import (
	"flag"
	"fmt"
)

func mian(){
	var name string
	flag.StringVar(&name,"name","傻逼","滚")
	flag.Parse()
	fmt.Printf("Hello, %v!\n",name)
}
