package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {

		src[i] = rand.Int63()
	}
}

func main(){
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for i := range lens {
		src := make([]int64, lens[i])
		fmt.Println(i)
		expect := make([]int64, lens[i])
		fmt.Println(expect)
		prepare(src)
		copy(expect, src)
		fmt.Println(copy(expect, src))
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		}
	}

