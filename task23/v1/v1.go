package main

import (
	"fmt"
	"github.com/infamax/L1/task23/errors"
)

/*
Удалить i-ый элемент из слайса.
*/

// DeleteItemsInSliceV1 первый способ без сохранения порядка
func DeleteItemsInSliceV1(nums []int, i int) ([]int, error) {
	if i >= len(nums) {
		return nil, errors.IncorrectPositionItem
	}

	if i < 0 {
		return nil, errors.NegativeIndex
	}

	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]

	nums = nums[:len(nums)-1]
	return nums, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(DeleteItemsInSliceV1(nums, 1))
}
