package main

import (
	"fmt"
	"reflect"
)

type Pet1 interface {
	Name() string
	Category() string
}

type Dog1 struct {
	name string // 名字。
}

func (dog *Dog1) SetName(name string) {
	dog.name = name
}

func (dog Dog1) Name() string {
	return dog.name
}

func (dog Dog1) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	var dog1 *Dog1
	fmt.Println("The first dog is nil.")
	dog2 := dog1
	fmt.Println("The second dog is nil.")
	var pet Pet1 = dog2
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n", dog2)
	fmt.Println()

	// 示例2。
	wrap := func(dog *Dog1) Pet1 {
		if dog == nil {
			fmt.Println("1")
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
}