package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// завершение горутины по контексту (функция отмены)
func worker1(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for{
        select{
        case <-ctx.Done():
            fmt.Println("worker1 finish")
            return
        default:
            someWork()
        }
    }
}

func example1() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker1(ctx, wg)
    
    go func(f context.CancelFunc){
        time.Sleep(time.Second)
        cancel()
    }(cancel)
    
    wg.Wait()
    fmt.Println("example1 finish")
}



// завершение горутины по контексту (Таймер)
func worker2(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for{
        select{
        case <-ctx.Done():
            fmt.Println("worker2 finish")
            return
        default:
            someWork()
        }
    }
}

func example2() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker2(ctx, wg)

    wg.Wait()
    fmt.Println("example2 finish")
}



// завершение горутины через чтение канала
func worker3(wg *sync.WaitGroup, finishCh <- chan struct{}) {
    defer wg.Done()

    for{
        select{
        case <-finishCh:
            fmt.Println("worker3 finish")
            return
        default:
            someWork()
        }
    }
}

func example3() {

    finishCh := make(chan struct{})

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker3(wg, finishCh)

    close(finishCh)
    wg.Wait()
    fmt.Println("example3 finish")
}



// завершение горутины при достижении ей результата
func worker4(wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i < 10; i++ {
        if i > 5 {
            fmt.Println("worker4 finish")
            return
        }
    }
}

func example4() {

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker4(wg)

    wg.Wait()
    fmt.Println("example4 finish")
}



// завершение горутины по таймеру
func worker5(out chan int) {
    someWork()
    someWork()
    out <- 42
}

func example5() {

    ch := make(chan int)
    go worker5(ch)

    select {
    case <-time.After(time.Second * 1):
        fmt.Println("timeout")
        break
    case <-ch:
        fmt.Println("worker5 finish")
        break
    }

    fmt.Println("example5 finish")
}



// завершение горутины по Goexit
func worker6(wg *sync.WaitGroup) {
    defer wg.Done()

    someWork()
    fmt.Println("worker6 finish")
    runtime.Goexit()
}

func example6() {

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker6(wg)

    wg.Wait()
    fmt.Println("example6 finish")
}



// завершение горутины по сигналу прерывания
func worker7(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()

    for{
        select{
        case <-ctx.Done():
            fmt.Println("worker7 finish")
            return
        default:
            someWork()
        }
    }
}

func example7() {

    ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
    defer cancel()

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker7(ctx, wg)

    select{ 
    case <-ctx.Done(): 
        fmt.Println("signal catch")
        break
    case <-time.After(time.Second * 3): 
        fmt.Println("timeout")
        break
    }

    cancel()
    wg.Wait()
    fmt.Println("example7 finish")
}



// выход по os.Exit
func worker8(wg *sync.WaitGroup) {
    defer wg.Done()

    someWork()
    os.Exit(0)
}

func example8() {

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go worker8(wg)

    wg.Wait()
    fmt.Println("example8 finish")
}


func main() {
    
    example1()
    example2()
    example3()
    example4()
    example5()
    example6()
    example7()
    example8()
}

func someWork() {
    time.Sleep(time.Second * 2)
}