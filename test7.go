package main

import "fmt"

func main(){
	numbers2 := []int{1, 2, 3, 4, 5, 6}

	maxIndex2 := len(numbers2) - 1

	for i, e := range numbers2 {
		fmt.Println(i,e)

		if i == maxIndex2 {

			numbers2[0] += e

		} else {

			numbers2[i+1] += e

		}

	}

	fmt.Println(numbers2)
}