package main

import (
	"flag"
	"time"
)

func main() {
	// Создаем канал для сигналов
	ch := make(chan struct{})
	// Используем флаг для указания количества воркеров
	n := flag.Int("workers", 1, "Number of workers to run")
	flag.Parse()
	// Запускаем n воркеров, которые будут слушать канал
	for i := range *n {
		go func(workerID int, ch chan struct{}) {
			println("Worker", workerID, "started")
			for {
				select {
				case <-ch:
					println("Worker", workerID, "received signal")
				default:
					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i, ch)
	}
	// Реализуем постоянную запись данных в канал
	for {
		ch <- struct{}{}
		time.Sleep(100 * time.Millisecond)
	}
}
