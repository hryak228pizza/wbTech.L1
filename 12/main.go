package main

import (
	"fmt"
)

func main() {

	s := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	result := []string{}

	for _, e := range s {
		if _, ok := m[e]; !ok {
			m[e] = struct{}{}
			result = append(result, e)
		}
	}

	fmt.Println(result)
}
