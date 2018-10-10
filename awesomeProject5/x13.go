package main

import "fmt"

func main() {

	s1 := make([]int, 5)//切片类型，长度
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)//切片类型，长度，容量（容量代表底层数组到底有多少元素，长度代表可以看到的元素）
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}
