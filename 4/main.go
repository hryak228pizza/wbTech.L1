package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, num int, wg *sync.WaitGroup) {

    defer wg.Done()
	fmt.Printf("Worker %v working...\n", num)

    for {
        select{
        case <-ctx.Done():
            return
        default:
            time.Sleep(2 * time.Second)
            fmt.Printf("Worker %v finished\n", num)
        }
    }
}

func main() {

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg := &sync.WaitGroup{}
	doneCh := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(shutdownCtx, i, wg)
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-doneCh:
		fmt.Println("program shutdown..")
		return
	case <-shutdownCtx.Done():
		fmt.Println("program shutdown..")
		return
	}
}