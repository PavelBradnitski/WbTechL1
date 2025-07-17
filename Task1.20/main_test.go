package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple sentence",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Sentence with leading and trailing spaces",
			input:    "  hello world  ",
			expected: "  world hello  ",
		},
		{
			name:     "Sentence with multiple spaces between words",
			input:    "a   b",
			expected: "b   a",
		},
		{
			name:     "Single word",
			input:    "word",
			expected: "word",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with only spaces",
			input:    "   ",
			expected: "   ",
		},
		{
			name:     "Sentence with punctuation",
			input:    "Hello, world!",
			expected: "world! Hello,",
		},
		{
			name:     "Mixed case",
			input:    "aBc dEf",
			expected: "dEf aBc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseWords(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %q, but got: %q", tc.expected, result)
			}
		})
	}
}
