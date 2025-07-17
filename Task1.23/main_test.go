package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteElem(t *testing.T) {
	testCases := []struct {
		name          string
		slice         []int
		index         int
		expectedSlice []int
		expectedError error
	}{
		{
			name:          "Valid delete at beginning",
			slice:         []int{1, 2, 3, 4, 5},
			index:         0,
			expectedSlice: []int{2, 3, 4, 5},
			expectedError: nil,
		},
		{
			name:          "Valid delete in the middle",
			slice:         []int{1, 2, 3, 4, 5},
			index:         2,
			expectedSlice: []int{1, 2, 4, 5},
			expectedError: nil,
		},
		{
			name:          "Valid delete at the end",
			slice:         []int{1, 2, 3, 4, 5},
			index:         4,
			expectedSlice: []int{1, 2, 3, 4},
			expectedError: nil,
		},
		{
			name:          "Empty slice",
			slice:         []int{},
			index:         0,
			expectedSlice: nil,
			expectedError: fmt.Errorf("index out of range"),
		},
		{
			name:          "Index out of range - negative",
			slice:         []int{1, 2, 3},
			index:         -1,
			expectedSlice: nil,
			expectedError: fmt.Errorf("index out of range"),
		},
		{
			name:          "Index out of range - too large",
			slice:         []int{1, 2, 3},
			index:         3,
			expectedSlice: nil,
			expectedError: fmt.Errorf("index out of range"),
		},
		{
			name:          "Slice with one element",
			slice:         []int{5},
			index:         0,
			expectedSlice: []int{},
			expectedError: nil,
		},
		{
			name:          "Slice with duplicate elements",
			slice:         []int{1, 2, 2, 3, 4},
			index:         2,
			expectedSlice: []int{1, 2, 3, 4},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultSlice, err := DeleteElem(tc.slice, tc.index)

			if tc.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error '%v', but got nil", tc.expectedError)
				}
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("Expected error '%v', but got '%v'", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(resultSlice, tc.expectedSlice) {
					t.Errorf("Expected slice: %v, but got: %v", tc.expectedSlice, resultSlice)
				}
			}
		})
	}
}
