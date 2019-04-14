package main

import (
	"errors"
	"fmt"
)

func echo(request string)(reponse string,err error){
	if request==""{
		err=errors.New("empty request")
		return
	}
	reponse=fmt.Sprintf("echo111:",request)
	return
}
func main()  {
	for  _,req:=range []string{"","shabi!"}{
		fmt.Printf("request",req)
		resp,err:=echo(req)
		if err!=nil{
			fmt.Println(err)
			continue
		}
		fmt.Println(resp)
	}
}
