package main

import "fmt"

type Beastf interface {
	F1()
	Init(app interface{})
}
type Base1 struct {
	App interface{}
}
type SonIntf interface {
	F2()
}
type Son struct {

}

func (s *Son)F2()  {

}
func (b *Base1)Init(app interface{})  {
	b.App=app
}
func (b *Base1)F1()  {
	v,ok:=((interface{})(b.App).(SonIntf))
	fmt.Println(v,ok)
	v.F2()
}
func main()  {

}
