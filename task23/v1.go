package main

import (
	"errors"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

var (
	incorrectPositionItem = errors.New("removing item number more than len slice")
	negativeIndex         = errors.New("negative index removing item")
)

// DeleteItemsInSliceV1 первый способ без сохранения порядка
func DeleteItemsInSliceV1(nums []int, i int) ([]int, error) {
	if i >= len(nums) {
		return nil, incorrectPositionItem
	}

	if i < 0 {
		return nil, negativeIndex
	}

	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]

	nums = nums[:len(nums)-1]
	return nums, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(DeleteItemsInSliceV1(nums, 1))
}
