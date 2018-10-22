package main

import "fmt"

type AnimaCaltegory struct {
	kingdom string
	phylum string
	class string
	order string
	family string
	genus string
	species string
}

func (ac AnimaCaltegory) String()string{
	return fmt.Sprintf(ac.class,ac.family,ac.species)
}
func main()  {
	a:=AnimaCaltegory{species:"shabi"}
	a.String()
	fmt.Println(a)
}