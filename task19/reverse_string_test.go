package main

import (
	"testing"
)

type test struct {
	name        string
	input       string
	expectedRes string
}

func TestReverseString(t *testing.T) {
	tests := []test{
		{
			name:        "empty string case",
			input:       "",
			expectedRes: "",
		},
		{
			name:        "len string = 1 case",
			input:       "a",
			expectedRes: "a",
		},
		{
			name:        "common case",
			input:       "главрыба",
			expectedRes: "абырвалг",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ReverseString(tt.input)
			if res != tt.expectedRes {
				t.Errorf("expectedRes = %v, but function return %v", tt.expectedRes, res)
			}
		})
	}
}
