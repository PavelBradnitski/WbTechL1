package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "String with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "Palindrome",
			input:    "madam",
			expected: "madam",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "String with special characters",
			input:    "!@#$%",
			expected: "%$#@!",
		},
		{
			name:     "String with Unicode characters",
			input:    "你好世界",
			expected: "界世好你",
		},
		{
			name:     "Mixed string",
			input:    "Hello, 世界!",
			expected: "!界世 ,olleH",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %q, but got: %q", tc.expected, result)
			}
		})
	}
}
