package main

import (
	"fmt"
	"sync"
	_"time"
)

var example []int

func initExample() {
	example = []int{2, 4, 6, 8, 10}
}

func square(n int, out chan int, wg *sync.WaitGroup) {
	out <- n*n
	wg.Done()
}

func main() {
	initExample()

	out := make(chan int)
	wg := &sync.WaitGroup{}

	go func() {
		for _, elem := range example {
			wg.Add(1)
			go square(elem, out, wg)
		}
	}()

	wg.Wait()
	for i := 1; i<= len(example); i++{
        fmt.Println(<- out)
    }
}



