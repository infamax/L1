package main

import (
	"reflect"
	"testing"
)

type test struct {
	name        string
	input       []int
	expectedRes []int
}

func TestQuickSort(t *testing.T) {
	tests := []test{
		{
			name:        "empty slice",
			input:       []int{},
			expectedRes: []int{},
		},
		{
			name:        "slice sorted decrease order",
			input:       []int{5, 4, 3, 2, 1},
			expectedRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:        "slice sorted increase order",
			input:       []int{1, 2, 3, 4, 5},
			expectedRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:        "slice elements in chaotic order",
			input:       []int{-7, -9, -11, 25, 11, 7, 12, 7},
			expectedRes: []int{-11, -9, -7, 7, 7, 11, 12, 25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.input, 0, len(tt.input)-1)
			if !reflect.DeepEqual(tt.input, tt.expectedRes) {
				t.Errorf("res = %v, but expectedRes = %v", tt.input, tt.expectedRes)
			}
		})
	}
}
