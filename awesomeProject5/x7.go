package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "oneｓ", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}
