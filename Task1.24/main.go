package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {
	point1 := Point{x: 0, y: 0}
	point2 := Point{x: 12, y: 5}

	distance := point1.Distance(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}
