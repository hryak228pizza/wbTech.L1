package main

import (
	"time"
    "context"
    "fmt"
    "sync"
)

const lifeTime = 1

func writer(ctx context.Context, c chan string, wg *sync.WaitGroup) {

    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Println("writer stop")
            close(c)
            return
        default:
            c <- "."
            time.Sleep(time.Second)
        }
    }    
}

func reader(c chan string, wg *sync.WaitGroup) {

    defer wg.Done()
    for v := range c {
        fmt.Println(v)
    }
    fmt.Println("reader stop")
}

func main() {

    ctx, cancel := context.WithTimeout(context.Background(), lifeTime * time.Second)
    defer cancel()

    c := make(chan string)
    wg := &sync.WaitGroup{}
    wg.Add(2)
    go writer(ctx, c, wg)
    go reader(c, wg)

    <-ctx.Done()
    fmt.Println("timer happened")
    wg.Wait()
}
