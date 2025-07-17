package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Empty nums1",
			nums1:    []int{},
			nums2:    []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty nums2",
			nums1:    []int{1, 2, 3},
			nums2:    []int{},
			expected: []int{},
		},
		{
			name:     "Both empty",
			nums1:    []int{},
			nums2:    []int{},
			expected: []int{},
		},
		{
			name:     "No intersection",
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Single intersection",
			nums1:    []int{1, 2, 3},
			nums2:    []int{3, 4, 5},
			expected: []int{3},
		},
		{
			name:     "Multiple intersections",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{3, 5, 6, 7},
			expected: []int{3, 5},
		},
		{
			name:     "All elements intersect",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums1 is a subset of nums2",
			nums1:    []int{1, 2},
			nums2:    []int{1, 2, 3},
			expected: []int{1, 2},
		},
		{
			name:     "nums2 is a subset of nums1",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Duplicates in nums1",
			nums1:    []int{1, 2, 2, 3},
			nums2:    []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Duplicates in nums2",
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Duplicates in both",
			nums1:    []int{1, 2, 2, 3},
			nums2:    []int{2, 2, 3, 4},
			expected: []int{2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := intersection(tc.nums1, tc.nums2)

			sort.Ints(result)
			sort.Ints(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
