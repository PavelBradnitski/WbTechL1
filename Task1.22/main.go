package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2 << 20)
	b := big.NewInt(2 << 21)

	// Сложение
	sum := new(big.Int).Add(a, b)

	// Вычитание
	diff := new(big.Int).Sub(a, b)

	// Умножение
	product := new(big.Int).Mul(a, b)

	// Деление
	quotient := new(big.Int).Div(b, a)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Сложение a + b =", sum)
	fmt.Println("Вычитание a - b =", diff)
	fmt.Println("Умножение a * b =", product)
	fmt.Println("Деление b / a = ", quotient)
}
