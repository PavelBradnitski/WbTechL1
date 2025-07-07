package main

import (
	"fmt"
	"sync"
)

// Использование sync.Mutex
type SyncMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		data: make(map[string]int),
	}
}

func (sm *SyncMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SyncMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	syncMap := NewSyncMap()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("%d", i)
			syncMap.Set(key, i*2)
			val, ok := syncMap.Get(key)
			if ok {
				fmt.Printf("Mutex: Key %s, Value %d\n", key, val)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Mutex: All goroutines finished.")

}
