package a

import (
	"fmt"
	"awesomeProject5/c"
)

type AAA struct {
	Name string
}

func (a AAA)GetName()string  {
	return a.Name
}
func (*AAA)DisAAA(obj interface{}) {
	fmt.Println(obj.(c.CCC).GetName())
}
