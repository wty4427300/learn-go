package main

import (
	"awesomeProject3/x2"
	"flag"
	"fmt"
	"os"
)
var name string
func inti(){
flag.StringVar(&name,"name","everyone","the greeting object")
//１．参数地址２．参数名３．参数默认值４．参数的简单说明
flag.CommandLine=flag.NewFlagSet("",flag.ExitOnError)
flag.CommandLine.Usage= func() {
	fmt.Fprintf(os.Stderr,"aaaaaa","bbbbbb")
	flag.PrintDefaults()
}
}
func main(){
	var a1 string
	fmt.Scan(&a1)
	fmt.Println(a1)
	flag.Parse()//用于解析命令行参数，并把它们的值赋给相应的变量
	fmt.Println("傻逼")
	fmt.Printf("Hello, %s!\n", name)
	x2.Main()
}
