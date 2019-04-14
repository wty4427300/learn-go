package main

import "fmt"

func main(){
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint8, uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Printf("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}

}
//byte等于uint8所以类型重复了。