package main

import (
	"testing"
)

type test struct {
	name        string
	offset      int
	val         int
	inputNum    int64
	expectedNum int64
	err         error
}

func TestChangeByteNum(t *testing.T) {
	tests := []test{
		{
			name:        "test case change to zero",
			offset:      1,
			val:         0,
			inputNum:    10,
			expectedNum: 8,
			err:         nil,
		},
		{
			name:        "test case change to one",
			offset:      2,
			val:         1,
			inputNum:    10,
			expectedNum: 14,
			err:         nil,
		},
		{
			name:        "error case",
			offset:      5,
			val:         1,
			inputNum:    10,
			expectedNum: 0,
			err:         incorrectPosition,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ChangeByteNum(tt.inputNum, tt.offset, tt.val)
			if err != tt.err {
				t.Errorf("error = %s, but expected error = %s", err, tt.err)
			}
			if res != tt.expectedNum {
				t.Errorf("res = %v, but expected res = %v", res, tt.expectedNum)
			}
		})
	}
}
