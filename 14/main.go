package main

import (
	"fmt"
)

func getType(v interface{}) {
	switch v.(type){
		case int: fmt.Println("int")
		case string: fmt.Println("string")
		case bool: fmt.Println("bool")
		case chan int: fmt.Println("chan int")
		default: fmt.Println("other")
	}
}

func main() {
	getType(1)
	getType("1")
	getType(true)
	getType(make(chan int))
}
