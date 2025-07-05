package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Используем контекст с таймаутом для управления временем выполнения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	// Запускаем горутину, которая будет слушать канал
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopping")
				return
			case <-ch:
				fmt.Println("Received signal")
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	// Основной цикл, который будет отправлять сигналы в канал
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("5 seconds passed, stopping")
			break loop
		default:
			// Отправка сигнала
			select {
			case ch <- 1:
				fmt.Println("Sent signal")
			case <-ctx.Done():
				break loop
			}

			// Пауза между отправками
			time.Sleep(1 * time.Second)
		}
	}
	// Дождаться завершения горутины без Sleep
	wg.Wait()
	fmt.Println("All goroutines finished")
}
