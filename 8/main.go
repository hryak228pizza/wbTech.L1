package main

import (
	"fmt"
)

func bitTo1(bit int64, bitPos int) int64 {
	bit |= 1<<(bitPos-1)
	return bit
} 

func bitTo0(bit int64, bitPos int) int64 {
	bit &^= 1<<(bitPos-1)
	return bit
} 

func main() {

	fmt.Println(bitTo0(5, 1)) // 101 -> 100 (5->4)
	fmt.Println(bitTo1(5, 2)) // 101 -> 111 (5->7)
    
}
