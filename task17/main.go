package main

import "fmt"

func BinarySearch(nums []int, val int) bool {
	l, r := -1, len(nums)

	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] < val {
			l = m
		} else {
			r = m
		}
	}

	if r == len(nums) {
		return false
	}

	if nums[r] != val {
		return false
	}

	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(BinarySearch(nums, 2))
	fmt.Println(BinarySearch(nums, -1))
	fmt.Println(BinarySearch(nums, 21))
}
