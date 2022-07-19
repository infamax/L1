package main

import "testing"

type test struct {
	name        string
	inputData   []int
	expectedRes int
}

func TestSumSquareNums(t *testing.T) {
	tests := []test{
		{
			name:        "empty slice",
			inputData:   []int{},
			expectedRes: 0,
		},
		{
			name:        "zero items",
			inputData:   []int{0, 0, 0},
			expectedRes: 0,
		},
		{
			name:        "positive items",
			inputData:   []int{2, 4, 6, 8},
			expectedRes: 120,
		},
		{
			name:        "negative items",
			inputData:   []int{-2, -4, -6, -8},
			expectedRes: 120,
		},
		{
			name:        "different numbers",
			inputData:   []int{2, -4, 6, -8},
			expectedRes: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SumSquareNums(tt.inputData)

			if res != tt.expectedRes {
				t.Errorf("SumSquareNums() = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
