package main

import (
	"fmt"
	"strconv"
	"sync"
)

type rating struct {
    ratingTable map[string]int
    mu sync.Mutex
}

func (r *rating) push(user string, points int) {
    r.mu.Lock()
    defer r.mu.Unlock()

    r.ratingTable[user] = points
}


func main() {
    
    r := &rating{
        ratingTable: make(map[string]int),
    }

    var wg sync.WaitGroup

    for i:=0; i < 10; i++ {
        wg.Add(1)
        go func(i int){
            defer wg.Done()
            r.push("user"+strconv.Itoa(i), i*10)
        }(i)
    }

    wg.Wait()
    fmt.Println(r.ratingTable)
}
