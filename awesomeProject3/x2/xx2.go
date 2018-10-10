package x2

import "flag"

var name string
func init(){
	flag.StringVar(&name,"name","傻逼xx2-1 ","没啥说明")
}
func Main()  {
	flag.Parse()
	hello(name)
}
