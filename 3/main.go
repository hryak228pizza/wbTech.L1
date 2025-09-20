package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Worker(c chan int) {
	for {		
		fmt.Println(<-c)
	}
}

func main() {
	args := os.Args
	poolSize, _ := strconv.Atoi(args[1])
	
	c := make(chan int)

	for i:=0; i < poolSize; i++{
		go Worker(c)
	}

	for {
		c <- time.Now().Second()
	}
}