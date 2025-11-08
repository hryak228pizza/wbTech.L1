package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5} // 1 4 9 16 25
	ch1 := make(chan int)       // х
	ch2 := make(chan int)       // х*х

	go func(arr []int, out chan<- int) {
		for _, elem := range arr {
			out <- elem
		}
		close(out)
	}(arr, ch1)

	go func(in <-chan int, out chan<- int) {
		for item := range in {
			out <- item * item
		}
		close(out)
	}(ch1, ch2)

	for item := range ch2 {
		fmt.Println(item)
	}
}
