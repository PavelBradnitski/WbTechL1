package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Println("Исходные значения:")
	fmt.Printf("x = %d, y = %d\n", x, y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Println("После обмена (сложение/вычитание):")
	fmt.Printf("x = %d, y = %d\n", x, y)

	// Сбрасываем значения для демонстрации XOR обмена.
	x = 1
	y = 2

	fmt.Println("Исходные значения (перед XOR обменом):")
	fmt.Printf("x = %d, y = %d\n", x, y)
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Println("После обмена (XOR):")
	fmt.Printf("x = %d, y = %d\n", x, y)
}
