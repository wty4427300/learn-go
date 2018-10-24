package main

type ax interface {
	aaa(int)int
}
type bx interface {
	bbb(int)int
}
type cx struct {
	id int
}
type dx struct {
	name string
	*cx
}

func (a *cx)bbb(i int)int  {
	return i
}
func (b *dx)aaa(i int)int  {
	return i
}
func main()  {
	var v1=new(cx)
	v1.id=1
	var v2=new(dx)
	v2.name="shapi"
	v2.cx=v1
}