package main

import (
	"reflect"
	"testing"
)

type test struct {
	name           string
	input1, input2 []int
	expectedRes    []int
}

func TestIntersection(t *testing.T) {
	tests := []test{
		{
			name:        "empty slices case",
			input1:      []int{},
			input2:      []int{},
			expectedRes: []int{},
		},
		{
			name:        "one empty slice case",
			input1:      []int{},
			input2:      []int{7, 1, 2, -1},
			expectedRes: []int{},
		},
		{
			name:        "no common elements in slice case",
			input1:      []int{1, 9, 6, 3, 2},
			input2:      []int{12, 4, 5},
			expectedRes: []int{},
		},
		{
			name:        "one common element case",
			input1:      []int{1, 2, 2, 1},
			input2:      []int{2, 2},
			expectedRes: []int{2},
		},
		{
			name:        "more than one common element in slices case",
			input1:      []int{4, 9, 5},
			input2:      []int{9, 4, 9, 8, 4},
			expectedRes: []int{4, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Intersection(tt.input1, tt.input2)
			if !reflect.DeepEqual(res, tt.expectedRes) {
				t.Errorf("res = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
