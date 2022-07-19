package main

import (
	"fmt"
	"github.com/infamax/L1/task23/errors"
)

// DeleteItemsInSliceV2 Порядок элементов важен
func DeleteItemsInSliceV2(nums []int, i int) ([]int, error) {
	if i >= len(nums) {
		return nil, errors.IncorrectPositionItem
	}

	if i < 0 {
		return nil, errors.NegativeIndex
	}

	for j := i; j < len(nums)-1; j++ {
		nums[j] = nums[j+1]
	}

	nums = nums[:len(nums)-1]
	return nums, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(DeleteItemsInSliceV2(nums, 2))
}
