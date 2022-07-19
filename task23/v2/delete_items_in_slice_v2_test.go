package main

import (
	"github.com/infamax/L1/task23/errors"
	"reflect"
	"testing"
)

type test struct {
	name          string
	inputData     []int
	i             int
	expectedRes   []int
	expectedError error
}

func TestDeleteItemsInSliceV2(t *testing.T) {
	tests := []test{
		{
			name:          "simple case",
			inputData:     []int{1, 2, 3, 4, 5},
			i:             2,
			expectedRes:   []int{1, 2, 4, 5},
			expectedError: nil,
		},
		{
			name:          "first item deleting",
			inputData:     []int{1, 2, 3, 4, 5},
			i:             0,
			expectedRes:   []int{2, 3, 4, 5},
			expectedError: nil,
		},
		{
			name:          "last item deleting",
			inputData:     []int{1, 2, 3, 4, 5},
			i:             4,
			expectedRes:   []int{1, 2, 3, 4},
			expectedError: nil,
		},
		{
			name:          "incorrect position removing item",
			inputData:     []int{1, 2, 3, 4, 5},
			i:             10,
			expectedRes:   nil,
			expectedError: errors.IncorrectPositionItem,
		},
		{
			name:          "negative index removing item",
			inputData:     []int{1, 2, 3, 4, 5},
			i:             -1,
			expectedRes:   nil,
			expectedError: errors.NegativeIndex,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DeleteItemsInSliceV2(tt.inputData, tt.i)
			if err != tt.expectedError {
				t.Errorf("err = %v, expectedRes = %v", err, tt.expectedError)
			}

			if !reflect.DeepEqual(res, tt.expectedRes) {
				t.Errorf("res = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
