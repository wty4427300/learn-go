package main

import (
	"awesomeProject5/a"
	"awesomeProject5/b"
)
func main()  {
	obj:=a.AAA{}
	obj1:=b.BBB{"BBB"}
	obj.DisAAA(&obj1)
	obj1.DisBBB(&obj)
}
