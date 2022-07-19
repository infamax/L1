package main

import "testing"

type test struct {
	name        string
	inputData   []int
	val         int
	expectedRes bool
}

func TestBinarySearch(t *testing.T) {
	tests := []test{
		{
			name:        "number exist in slice",
			inputData:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			val:         2,
			expectedRes: true,
		},
		{
			name:        "number more than max number in slice",
			inputData:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			val:         21,
			expectedRes: false,
		},
		{
			name:        "number less than min number in slice",
			inputData:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			val:         -1,
			expectedRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := BinarySearch(tt.inputData, tt.val)
			if res != tt.expectedRes {
				t.Errorf("res = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
