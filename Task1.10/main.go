package main

import (
	"fmt"
	"math"
)

func main() {
	var temperatures = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)
	var keyValue int
	for _, v := range temperatures {
		if v > 0 {
			keyValue = int(math.Floor(v/10) * 10)
		} else {
			keyValue = int(math.Ceil(v/10) * 10)
		}
		m[keyValue] = append(m[keyValue], v)
	}
	fmt.Println(m)
}
