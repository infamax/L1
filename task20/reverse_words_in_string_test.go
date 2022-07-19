package main

import (
	"testing"
)

type test struct {
	name        string
	input       string
	expectedStr string
}

func TestReverseWordsInStrings(t *testing.T) {
	tests := []test{
		{
			name:        "empty string case",
			input:       "",
			expectedStr: "",
		},
		{
			name:        "simple case",
			input:       "snow dog sun",
			expectedStr: "sun dog snow",
		},
		{
			name:        "one word case",
			input:       "snow",
			expectedStr: "snow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ReverseWordsInStrings(tt.input)
			if res != tt.expectedStr {
				t.Errorf("expectedString = %v, but function return = %v", tt.expectedStr, res)
			}
		})
	}
}
