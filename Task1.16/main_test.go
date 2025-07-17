package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicate elements",
			input:    []int{5, 2, 8, 1, 9, 2, 4},
			expected: []int{1, 2, 2, 4, 5, 8, 9},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-5, 2, -8, 1, 9, -2, 4},
			expected: []int{-8, -5, -2, 1, 2, 4, 9},
		},
		{
			name:     "Array with mixed positive and negative numbers",
			input:    []int{-1, 0, 1, -2, 2},
			expected: []int{-2, -1, 0, 1, 2},
		},
		{
			name:     "Random array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			result := quickSort(inputCopy)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
