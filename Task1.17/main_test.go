package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Target found in the middle",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   11,
			expected: 4,
		},
		{
			name:     "Target found at the beginning",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   2,
			expected: 0,
		},
		{
			name:     "Target found at the end",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   12,
			expected: 5,
		},
		{
			name:     "Target not found - less than smallest",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   1,
			expected: -1,
		},
		{
			name:     "Target not found - greater than largest",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   13,
			expected: -1,
		},
		{
			name:     "Target not found - in the middle",
			arr:      []int{2, 5, 7, 8, 11, 12},
			target:   9,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element array - target found",
			arr:      []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element array - target not found",
			arr:      []int{5},
			target:   10,
			expected: -1,
		},
		{
			name:     "Array with duplicate elements - target found",
			arr:      []int{2, 5, 5, 7, 8, 11, 12},
			target:   5,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.arr, tc.target)
			if result != tc.expected {
				t.Errorf("Expected: %d, but got: %d", tc.expected, result)
			}
		})
	}
}
