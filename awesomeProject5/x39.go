package main

import "fmt"

func main(){
	var halffun= map[string]func(int){
		"aaaaa": func(i int) {
			fmt.Println("shabi", i)
		},
		"bbbbb": func(i int) {
			fmt.Println("gundan", i)
		},
	}
	halffun["aaaaa"](1)
}
