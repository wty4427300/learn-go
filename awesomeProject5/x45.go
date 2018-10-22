package main

import "fmt"

func  modifyArray(a [3]string)[3]string  {
	a[1]="x"
	return a
}
func main(){
	array1:=[3]string{"a","b","c"}
	fmt.Println(array1)
	array2:=modifyArray(array1)
	fmt.Println(array1)
	fmt.Println(array2)
}
