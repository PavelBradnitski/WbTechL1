package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Используем context.WithCancel из за того что он легко имплементируется: контекст можно передавать во все воркеры и производные функции
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Обработка сигнала Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Println("Received interrupt signal. Cancelling context...")
		cancel()
	}()

	n := flag.Int("workers", 1, "Number of workers to run")
	flag.Parse()
	for i := range *n {
		wg.Add(1)
		go func(workerID int) {
			go worker(ctx, i, &wg)
		}(i)
	}

	fmt.Println("The program is running. Press Ctrl+C to exit.")
loop:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Main loop: context cancelled, stopping...")
			break loop
		default:
			fmt.Printf("Main loop iteration %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}
	// Ждём завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers stopped. Exiting.")
}
