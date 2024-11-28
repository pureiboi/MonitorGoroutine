package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func monitorGoroutines(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is done")
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
		}
	}
}

func worker(wg *sync.WaitGroup, id int, ch <-chan int, done chan<- int) {

	defer wg.Done()
	for x := range ch {
		time.Sleep(time.Second * 2)
		fmt.Printf("worker %d display x from channel %d\n", id, x)
	}
	done <- 1
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	noOfWorker := 5
	ch := make(chan int, 100)
	done := make(chan int, 100)

	wg.Add(1)
	go monitorGoroutines(ctx, wg)

	for x := range noOfWorker {
		wg.Add(1)
		go worker(wg, x, ch, done)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			if noOfWorker == len(done) {
				cancel()
				fmt.Println("context is cancelled")
				break
			}
		}

	}()

	for x := range 100 {
		ch <- x
	}
	close(ch)

	wg.Wait()
	fmt.Println("All operation finished")
}
