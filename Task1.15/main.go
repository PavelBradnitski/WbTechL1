package main

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// проблема могла возникнуть если длина v окажется < 100
	if len(v) > 100 {
		justString = v[:100]
	} else {
		justString = v
	}

}

func createHugeString(i int) string {
	return strings.Repeat("A", i)
}

func main() {
	someFunc()
}
