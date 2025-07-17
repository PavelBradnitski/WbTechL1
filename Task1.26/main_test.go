package main

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "All unique lowercase",
			input:    "abcd",
			expected: true,
		},
		{
			name:     "All unique uppercase",
			input:    "ABCD",
			expected: true,
		},
		{
			name:     "Mixed case with duplicates",
			input:    "abCdefAaf",
			expected: false,
		},
		{
			name:     "Lowercase with duplicates",
			input:    "aabcd",
			expected: false,
		},
		{
			name:     "Unique Cyrillic",
			input:    "аБвГ",
			expected: true,
		},
		{
			name:     "Cyrillic with duplicates",
			input:    "aВвГа",
			expected: false,
		},
		{
			name:     "Space",
			input:    " ",
			expected: true,
		},
		{
			name:     "Repeated character",
			input:    "aa",
			expected: false,
		},
		{
			name:     "Long unique string",
			input:    "asdfghjkl",
			expected: true,
		},
		{
			name:     "Mixed case",
			input:    "AbCdEf",
			expected: true,
		},
		{
			name:     "Numbers",
			input:    "1234567890",
			expected: true,
		},
		{
			name:     "Symbols",
			input:    "!@#$%^&*()",
			expected: true,
		},
		{
			name:     "Mixed numbers and letters",
			input:    "a1b2c3d4",
			expected: true,
		},
		{
			name:     "Mixed numbers and letters with duplicates",
			input:    "a1b2a3d4",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isUnique(tc.input)
			if result != tc.expected {
				t.Errorf("For input '%s', expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
