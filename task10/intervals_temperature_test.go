package main

import (
	"reflect"
	"sort"
	"testing"
)

type test struct {
	name        string
	temps       []float64
	expectedRes map[int][]float64
}

func TestIntervalsTemperature(t *testing.T) {
	tests := []test{
		{
			name:  "one interval case",
			temps: []float64{12.3, 11.2, 19.7, 13.4, 15.5},
			expectedRes: map[int][]float64{
				10: {11.2, 12.3, 13.4, 15.5, 19.7},
			},
		},
		{
			name:  "one value in each interval case",
			temps: []float64{12.3, 31.9, -20.5},
			expectedRes: map[int][]float64{
				-20: {-20.5},
				10:  {12.3},
				30:  {31.9},
			},
		},
		{
			name:  "basic case",
			temps: []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			expectedRes: map[int][]float64{
				-20: {-25.4, -27.0, -21.0},
				10:  {13.0, 19.0, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IntervalsTemperature(tt.temps)
			for _, val := range res {
				sort.Float64s(val)
			}
			if reflect.DeepEqual(res, tt.expectedRes) {
				t.Errorf("function return: %v, expectedRes: %v", res, tt.expectedRes)
			}
		})
	}
}
