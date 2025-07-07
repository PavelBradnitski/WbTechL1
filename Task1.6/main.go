package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Выход по условию
	fmt.Println("Starting example with exit by condition.")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Working:")
			time.Sleep(time.Second)
			if i == 2 {
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("Example with exit by condition finished.")
	// Выход по каналу уведомления
	fmt.Println("Starting example with notification channel.")
	done := make(chan bool)

	go func(done chan bool) {
		for {
			select {
			case <-done:
				fmt.Println("Received signal to stop.")
				return
			default:
				fmt.Println("Working:")
				time.Sleep(time.Second)
			}
		}
	}(done)

	time.Sleep(2 * time.Second)

	fmt.Println("Sending signal to stop worker.")
	done <- true

	time.Sleep(time.Second)
	fmt.Println("Example with notification channel finished.")
	// Выход по контекст с таймаутом
	fmt.Println("Starting example with context.")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopping")
				return
			default:
				fmt.Println("Working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	wg.Wait()
	fmt.Println("Example with context finished.")
	// Выход по runtime.Goexit
	fmt.Println("Starting example with runtime.Goexit.")
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Worker: Defer called before exit.")

		for i := 0; i < 10; i++ {
			fmt.Println("Working:", i)
			time.Sleep(time.Second)

			if i == 4 {
				fmt.Println("Worker: Exiting using runtime.Goexit().")
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
	fmt.Println("Example with runtime.Goexit finished.")
}
