package main

import (
	"fmt"
)

func main() {

	A := []int{1, 2, 3, 2}
	B := []int{2, 3, 2}

	intersect := make(map[int]int)
	for _, elem := range A {
		intersect[elem] ++
	}
	for _, elem := range B {
		intersect[elem] ++
	}

	for k, v := range intersect {
		if v > 1 {
			fmt.Println(k)
		}
	}
}
