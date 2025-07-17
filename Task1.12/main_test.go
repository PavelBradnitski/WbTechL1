package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Slice with single unique string",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "Slice with all same strings",
			input:    []string{"hello", "hello", "hello"},
			expected: []string{"hello"},
		},
		{
			name:     "Slice with multiple unique strings",
			input:    []string{"hello", "world", "hello", "go", "world"},
			expected: []string{"go", "hello", "world"}, // CORRECT ORDER
		},
		{
			name:     "Slice with empty strings",
			input:    []string{"", "", "hello", "", "world"},
			expected: []string{"", "hello", "world"}, // CORRECT ORDER
		},
		{
			name:     "Slice with different casing",
			input:    []string{"Hello", "hello", "HELLO"},
			expected: []string{"HELLO", "Hello", "hello"}, // CORRECT ORDER
		},
		{
			name:     "Slice with mixed case and empty strings",
			input:    []string{"", "Hello", "", "hello"},
			expected: []string{"", "Hello", "hello"}, // CORRECT ORDER
		},
		{
			name:     "Slice with numbers as strings",
			input:    []string{"1", "2", "1", "3"},
			expected: []string{"1", "2", "3"}, // CORRECT ORDER
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uniqueStrings(tc.input)

			// Sort both result and expected so order doesn't matter for comparison
			sort.Strings(result)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
