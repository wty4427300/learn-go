package main

import "fmt"

var container3 = []string{"zero", "one", "two"}

func main() {
	var ok bool
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
	value, ok := interface{}(container).([]string)
	fmt.Println(value,ok)
}
