package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type person123 struct {
	int
	name string
}

func (p person123)Dis()  {
	fmt.Println(p)
}

type student struct {
	person123
	age int
	Sex int
}
func (s student)Sayhello(toname string)(string,int){ //如果方法不是大写第一个字母的公开方法，是无法被反射获取的
	return s.name+" say hello to "+toname,1
}
func (s *student)Dis(){ //*student不属于student 仅属于*student对象
	fmt.Println(s)
}
func main()  {
	s:=student{person123:person123{int:9,name:"shapi"},age:22,Sex:0}
	t:=reflect.TypeOf(s)
	v:=reflect.ValueOf(s)
	fmt.Println(t)
	fmt.Println(v)

	for i:=0; i<t.NumField(); i++ {
		key:=t.Field(i)
		val:=t.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}
	fmt.Println(t.Kind())
	for i:=0;i<t.NumMethod();i++{  //仅能获取非指针对象的引用方法
		m:=t.Method(i)
		fmt.Println(m.Name,m.Type)
	}
	t1:=reflect.TypeOf(&s)
	v1:=reflect.ValueOf(&s)
	fmt.Println(t1)
	fmt.Println(v1)
	if k:=t1.Kind();k==reflect.Ptr{
		tt:=t1.Elem() //指针通过elem获取内容
		vv:=v1.Elem()
		for i:=0;i<tt.NumField();i++{
			key:=tt.Field(i)
			val:=vv.Field(i)
			fmt.Println(key.Name,key.Type,"|",val)
		}
	}
	fmt.Println(t.FieldByName("person123"))
	fmt.Println(t.FieldByIndex([]int{0}))//取到类
	fmt.Println(t.FieldByIndex([]int{0,0}),t.FieldByIndex([]int{0,1}))//取到类的属性
	m2, _ :=t.MethodByName("Sayhello")
	fmt.Println(m2)
	x:=123
	vx:=reflect.ValueOf(&x)
	fmt.Println(vx)
	vx.Elem().SetInt(333)
	fmt.Println(x)
	fmt.Println("===========================")
	s11:=student{person123:person123{int:1,name:"aaa"},age:22,Sex:1}
	fmt.Println(s11)
	v11:=reflect.ValueOf(&s11)
	v11.Elem().FieldByName("age").CanSet()//小写是只读
	fmt.Println(v11.Elem().FieldByName("Sex").CanSet())//大写才是读写
	if v11.Elem().FieldByName("Sex").CanSet(){
		v11.Elem().FieldByName("Sex").SetInt(2)
	}
	fmt.Println("==================")
	fv := reflect.ValueOf(prints)//获取一个普通函数的句柄
    fmt.Println(&fv)
	params := make([]reflect.Value,1)  //参数
	params[0] = reflect.ValueOf(20)   //参数设置为20
	rs := fv.Call(params) //rs作为结果接受函数的返回值
	fmt.Println(rs[0])
	fmt.Println("result: "+rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()
}
func prints(i int) string {
	fmt.Println("i =",i)
	return strconv.Itoa(i)
}