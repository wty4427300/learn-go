package main

import "fmt"

func main(){
	nums:=[]int{2,3,4}
	sum:=0
	for _,num:=range nums{
		sum +=num

	}
	fmt.Println("sum:",sum)
	for i, num:=range nums{
		if num==3{
			fmt.Println(i)
		}
	}
	kvs:=map[string]string{"a":"apple","b":"banana"}
	for k,v:=range kvs {
		fmt.Println(k+":1",v+":2")
	}
	for i,c:=range "gooooooooaaaa0000"{
		fmt.Println(i,c)
	}
}