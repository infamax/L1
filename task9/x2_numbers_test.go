package main

import (
	"reflect"
	"sort"
	"testing"
)

type test struct {
	name        string
	inputDate   []int
	expectedRes []int
}

func TestGetNumsChan(t *testing.T) {
	tests := []test{
		{
			name:        "empty slice",
			inputDate:   []int{},
			expectedRes: []int{},
		},
		{
			name:        "zeroes items",
			inputDate:   []int{0, 0, 0, 0, 0},
			expectedRes: []int{0, 0, 0, 0, 0},
		},
		{
			name:        "positive items",
			inputDate:   []int{2, 4, 6, 8},
			expectedRes: []int{4, 8, 12, 16},
		},
		{
			name:        "negative items",
			inputDate:   []int{-2, -4, -6, -8},
			expectedRes: []int{-4, -8, -12, -16},
		},
		{
			name:        "different items",
			inputDate:   []int{-2, 4, -6, 8},
			expectedRes: []int{-4, 8, -12, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := make([]int, 0)

			for val := range X2Nums(GetNumsChan(tt.inputDate)) {
				res = append(res, val)
			}

			sort.Ints(res)
			sort.Ints(tt.expectedRes)

			if !reflect.DeepEqual(res, tt.expectedRes) {
				t.Errorf("X2nums = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
