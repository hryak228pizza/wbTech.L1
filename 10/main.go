package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, elem := range arr {
		key := math.Trunc(elem/10) * 10
		groups[int(key)] = append(groups[int(key)], elem)
	}

	fmt.Println(groups)
}
