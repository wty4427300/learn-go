package b

import (
	"fmt"
	"awesomeProject5/c"
)

type BBB struct {
	Name string
}
func (a BBB)GetName() string{
	return a.Name
}
//func (*BBB)DisBBB(obj * a.AAA) {
//	fmt.Println(obj.Name)
//}

func (*BBB)DisBBB(obj interface{}) {
	fmt.Println(obj.(c.CCC).GetName())
}