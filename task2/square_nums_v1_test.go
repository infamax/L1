package main

import (
	"reflect"
	"sort"
	"testing"
)

type test struct {
	name      string
	testSlice []int
	want      []int
}

func TestSquareNumsV1(t *testing.T) {
	tests := []test{
		{
			name:      "empty slice",
			testSlice: []int{},
			want:      []int{},
		},
		{
			name:      "positive numbers",
			testSlice: []int{2, 4, 6, 8},
			want:      []int{4, 16, 36, 64},
		},
		{
			name:      "negative numbers",
			testSlice: []int{-2, -4, -6, -8},
			want:      []int{4, 16, 36, 64},
		},
		{
			name:      "zero numbers",
			testSlice: []int{0, 0, 0},
			want:      []int{0, 0, 0},
		},
		{
			name:      "different numbers",
			testSlice: []int{-1, 1, 2, 5, -7},
			want:      []int{1, 1, 4, 25, 49},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := SquareNumsV1(tt.testSlice)
			res := make([]int, 0, len(ch))

			for val := range ch {
				res = append(res, val)
			}

			sort.Ints(tt.want)
			sort.Ints(res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("SquareNums() = %v, want = %v", res, tt.want)
			}
		})
	}
}
