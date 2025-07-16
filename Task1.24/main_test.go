package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		p1       Point
		p2       Point
		expected float64
	}{
		{NewPoint(0, 0), NewPoint(0, 0), 0},
		{NewPoint(0, 0), NewPoint(3, 4), 5},
		{NewPoint(1, 1), NewPoint(4, 5), 5},
		{NewPoint(-1, -1), NewPoint(-4, -5), 5},
		{NewPoint(2.5, 3.5), NewPoint(2.5, 3.5), 0},
	}

	for _, tt := range tests {
		actual := tt.p1.Distance(tt.p2)
		if actual != tt.expected {
			t.Errorf("Distance between %v and %v: got %v, expected %v",
				tt.p1, tt.p2, actual, tt.expected)
		}
	}
}
