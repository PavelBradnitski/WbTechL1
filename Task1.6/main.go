package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting example with exit by condition.")
	var wg sync.WaitGroup
	wg.Add(1)
	// Выход по условию
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
	fmt.Println("Starting example with notification channel.")
	// канал уведомления
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
	fmt.Println("Starting example with context.")

	// Используем контекст с таймаутом для управления временем выполнения
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Add(1)
	// Запускаем горутину, которая будет слушать канал
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
