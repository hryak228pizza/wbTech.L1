package main

import (
	"fmt"
)

func main() {

	x := 2
	y := 3

	// 2 3
	// 5 3
	// 5 2
	// 3 2

	// 010 101
	// 111 101
	// 111 010
	// 101 010

	fmt.Printf("Initial values: x:%v, y:%v\n", x, y)

	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("values after +/-: x:%v, y:%v\n", x, y)

	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Printf("values after XOR: x:%v, y:%v\n", x, y)
}
