package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name()string
	Cattegory()string
}

type Dog struct {
	name string
}
func(dog *Dog)SetName(name string){
	dog.name=name
}
func (dog Dog)Name()string  {
	return dog.name
}
func (dog Dog)Cattegory()string  {
	return "dog"
}
func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Cattegory(), pet.Name())
}