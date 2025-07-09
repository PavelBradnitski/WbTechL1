package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Count int
	mu    sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.mu.Lock()
			c.Count++
			fmt.Printf("Значение равно %v\n", c.Count)
			c.mu.Unlock()
		}()
	}
	wg.Wait()
}
