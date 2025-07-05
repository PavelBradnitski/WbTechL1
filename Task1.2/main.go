package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	numArray := []int{2, 4, 6, 8, 10}
	// Используем для ожидания выполнения всех горутин
	var wg sync.WaitGroup
	for _, v := range numArray {
		wg.Add(1)
		// работает на версии go >= 1.22
		go func() {
			Square(v)
			wg.Done()
		}()
		// Если версия go < 1.22 используем данную конструкцию
		// go func(n int) {
		// 	Square(n)
		// 	wg.Done()
		// }(v)
	}
	// Ожидаем выполнения всех горутин
	wg.Wait()
}

func Square(n int) {
	fmt.Printf("Square of %d is %d\n", n, int(math.Pow(float64(n), 2)))
}
