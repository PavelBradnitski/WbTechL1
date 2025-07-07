package main

import (
	"fmt"
	"sync"
)

// 1. Использование sync.Mutex
type SafeMapMutex struct {
	data map[string]int
	mu   sync.Mutex
}

func NewSafeMapMutex() *SafeMapMutex {
	return &SafeMapMutex{
		data: make(map[string]int),
	}
}

func (sm *SafeMapMutex) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMapMutex) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	safeMapMutex := NewSafeMapMutex()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("%d", i)
			safeMapMutex.Set(key, i*2)
			val, ok := safeMapMutex.Get(key)
			if ok {
				fmt.Printf("Mutex: Key %s, Value %d\n", key, val)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Mutex: All goroutines finished.")

}
