package main

import "fmt"

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	s4=s4[0:cap(s4)]
	s4=s4[0:2]
	s3 =[]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println(s3,cap(s3))
}