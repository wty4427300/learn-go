package main

func main(){
	a:=1
	b:=2
	c1:=[]int{1,2}
	yyy(c1,a,b)
	cc:=[]int{3,4}
	yyy(c1,cc...)
}
func yyy(xx2 []int,a... int){
  a[0]=123
	c1 := make([]int,1)
	c1[0]=123
}