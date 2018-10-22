package main

import "fmt"

type Stu struct {
	name string
	age int
	score int
	class string
}
type Tea struct {
	name string
	age int
	class string
}
func (stu Stu) bb (tea Tea)(string,int,string) {
	return tea.name,tea.age,tea.class
}
func (tea Tea) aa (stu Stu)(string,int,int,string) {
	return stu.name,stu.age,stu.score,stu.class
}

func main(){
	var stu1  = Stu{"ergou",8,12,"erban"}
	var tea1  =Tea{"xixi",23,"erban"}
	name,age,class:=stu1.bb(tea1)
	name1,age1,score1,class1:=tea1.aa(stu1)
	fmt.Println(stu1.name,"tea:",name,age,class)
	fmt.Println(tea1.name,"stu:",name1,age1,score1,class1)
}
