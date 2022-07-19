package main

import "testing"

type test struct {
	name        string
	i           interface{}
	expectedRes string
}

func TestShowMeTheType(t *testing.T) {
	tests := []test{
		{
			name:        "int case ",
			i:           1,
			expectedRes: "int",
		},
		{
			name:        "[]int case",
			i:           []int{1, 2, 3},
			expectedRes: "[]int",
		},
		{
			name:        "string case",
			i:           "abc",
			expectedRes: "string",
		},
		{
			name:        "bool case",
			i:           false,
			expectedRes: "bool",
		},
		{
			name:        "float64 case",
			i:           2.5,
			expectedRes: "float64",
		},
		{
			name:        "chan int case",
			i:           make(chan int),
			expectedRes: "chan int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ShowMeTheType(tt.i)
			if res != tt.expectedRes {
				t.Errorf("expectedRes = %v, but function return = %v", tt.expectedRes, res)
			}
		})
	}
}
