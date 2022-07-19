package main

import "testing"

type test struct {
	name        string
	input       string
	expectedRes bool
}

func TestIsUniqueString(t *testing.T) {
	tests := []test{
		{
			name:        "empty string case",
			input:       "",
			expectedRes: true,
		},
		{
			name:        "unique string case",
			input:       "abcd",
			expectedRes: true,
		},
		{
			name:        "string with letters different registers case",
			input:       "abCdefAa",
			expectedRes: false,
		},
		{
			name:        "string repeated chars case",
			input:       "aabcd",
			expectedRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsUniqueString(tt.input) != tt.expectedRes {
				t.Errorf("expectedRes = %v, but function result = %v", tt.expectedRes, IsUniqueString(tt.input))
			}
		})
	}
}
