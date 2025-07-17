package main

import (
	"errors"
	"testing"
)

func TestSetBit(t *testing.T) {
	testCases := []struct {
		name        string
		num         int64
		i           uint
		bitValue    int
		expectedNum int64
		expectedErr error
	}{
		{
			name:        "Set bit 0 to 1",
			num:         0,
			i:           0,
			bitValue:    1,
			expectedNum: 1,
			expectedErr: nil,
		},
		{
			name:        "Set bit 0 to 0",
			num:         1,
			i:           0,
			bitValue:    0,
			expectedNum: 0,
			expectedErr: nil,
		},
		{
			name:        "Set bit 1 to 1",
			num:         0,
			i:           1,
			bitValue:    1,
			expectedNum: 2,
			expectedErr: nil,
		},
		{
			name:        "Set bit 1 to 0",
			num:         2,
			i:           1,
			bitValue:    0,
			expectedNum: 0,
			expectedErr: nil,
		},
		{
			name:        "Set bit 2 to 1",
			num:         5,
			i:           2,
			bitValue:    1,
			expectedNum: 5,
			expectedErr: nil,
		},
		{
			name:        "Set bit 2 to 0",
			num:         5,
			i:           2,
			bitValue:    0,
			expectedNum: 1,
			expectedErr: nil,
		},
		{
			name:        "Invalid bitValue (2)",
			num:         0,
			i:           0,
			bitValue:    2,
			expectedNum: 0,
			expectedErr: errors.New("bitValue must be 0 or 1"),
		},
		{
			name:        "Invalid bitValue (-1)",
			num:         0,
			i:           0,
			bitValue:    -1,
			expectedNum: 0,
			expectedErr: errors.New("bitValue must be 0 or 1"),
		},
		{
			name:        "Set bit with larger number",
			num:         123456789,
			i:           5,
			bitValue:    1,
			expectedNum: 123456821,
			expectedErr: nil,
		},

		{
			name:        "Set bit with larger number to 0",
			num:         123456821,
			i:           5,
			bitValue:    0,
			expectedNum: 123456789,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SetBit(tc.num, tc.i, tc.bitValue)

			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("Expected error: %v, but got nil", tc.expectedErr)
				}
				if err.Error() != tc.expectedErr.Error() {
					t.Fatalf("Expected error: %v, but got: %v", tc.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}

			if result != tc.expectedNum {
				t.Errorf("Expected result: %d, but got: %d", tc.expectedNum, result)
			}
		})
	}
}
