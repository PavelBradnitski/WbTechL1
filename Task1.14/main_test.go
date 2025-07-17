package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestDefineType(t *testing.T) {
	testCases := []struct {
		name           string
		input          interface{}
		expectedOutput string // Hard to test the print statement directly, so use a simple string check
		expectedPanic  bool   // Add panic check because testing Printf directly is difficult
	}{
		{
			name:           "String",
			input:          "hello",
			expectedOutput: "string",
		},
		{
			name:           "Integer",
			input:          123,
			expectedOutput: "int",
		},
		{
			name:           "Boolean",
			input:          true,
			expectedOutput: "boolean",
		},
		{
			name:           "Channel",
			input:          make(chan int),
			expectedOutput: "chan",
		},
		{
			name:           "Float",
			input:          3.14,
			expectedOutput: "Unknown type",
		},
		{
			name:           "Nil",
			input:          nil,
			expectedOutput: "Unknown type",
		},
		{
			name:           "Struct",
			input:          struct{}{},
			expectedOutput: "Unknown type",
		},
		{
			name:           "Slice",
			input:          []int{1, 2, 3},
			expectedOutput: "Unknown type",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Capture the output of DefineType.  A more robust approach might involve
			// a custom writer instead of capturing stdout, but this is sufficient for basic testing.
			originalStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe: %v", err)
			}
			os.Stdout = w

			defer func() {
				os.Stdout = originalStdout
				w.Close()
				r.Close()
			}()

			// Handle panics. If test case specifies a panic, recover and verify.
			if tc.expectedPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic as expected")
					}
				}()
			}

			DefineType(tc.input) // Call the function being tested

			w.Close()
			out, err := io.ReadAll(r)
			if err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}

			output := string(out)

			// Check if the output contains the expected string
			if !strings.Contains(output, tc.expectedOutput) {
				t.Errorf("For input %v, expected output to contain '%s', but got '%s'", tc.input, tc.expectedOutput, output)
			}

		})
	}
}
