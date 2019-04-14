// merge Sort
package main

import "fmt"

func mergeSort(c []int) []int {
	var n = len(c)
	if n < 2 {
		return c
	}
	var m = n / 2
	var lCol = c[:m]
	fmt.Println(lCol)
	var lSize = len(lCol)
	if lSize >= 2 {
		lCol = mergeSort(lCol)
	}

	var rCol = c[m:]
	var rSize = len(rCol)
	if rSize >= 2 {
		rCol = mergeSort(rCol)
	}
	var tmp = make([]int, 0)
	var i, j = 0, 0
	for i < lSize && j < rSize {
		if lCol[i] < rCol[j] {
			tmp = append(tmp, lCol[i])
			i++
		} else {
			tmp = append(tmp, rCol[j])
			j++
		}

	}
	if i < lSize {
		tmp = append(tmp, lCol[i:]...)
	}
	if j < rSize {
		tmp = append(tmp, rCol[j:]...)
	}
	return tmp
}
func main() {
	var b = []int{35, 9, 8, 6, 20, 7, 4, 3, 2, 1}

	fmt.Println(mergeSort(b))
}