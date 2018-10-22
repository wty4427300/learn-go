package main

import "fmt"

type q1 struct {
  id int
}

func (self *q1)myQ1(c int,d int) int{
	return  c+d
}
func (self  q1)myQ2(c int,d int) int{
	return  c+d
}
func main(){
	var q2=new(q1)
	fmt.Println(q2.myQ1(1,2))
	var q3=q1{1}
	fmt.Println(q3.myQ2(2,3))

}
