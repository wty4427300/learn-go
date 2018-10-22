package main

import (
	"fmt"
)

func main()  {
	type person struct {
		id int
		name string
	}
	var p1=person{1,"shabi"}
	var p2=person{}
	p2.name="gundan"
	p2.id=2
	fmt.Println(p1,p2)
	var p3= struct {
		id int
		name string
		style struct{
			color string
			model int
		}
	}{3,"shapi", struct {
		color string
		model int
	}{color: "hongse", model:2}}
	fmt.Println(p3)
	var p4=new(person)
	p4.id=1
	p4.name="aaaaaaaaa"
	(*p4).id=1
	(*p4).name="aaaaaaaaa"
	fmt.Println(p4,*p4)
    fmt.Println("===========================")
	type Person struct {
		id int
		name string
	}
	type stu struct {
		Person
		age int
	}
	type tea struct {
		Person
		die int
	}
	var s1=stu{
		Person:Person{id:1,name:"aaa"},
		age:23,
	}
	var t1=tea{}
	t1.Person.id=2
	t1.id=2
	t1.Person.name="bbbb"
	t1.name="bbbb"
	t1.die=54
	fmt.Println(s1,t1)
}
