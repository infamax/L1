package main

import "fmt"

func QuickSort(nums []int, l, r int) {
	if l < r {
		q := Partition(nums, l, r)
		QuickSort(nums, l, q)
		QuickSort(nums, q+1, r)
	}
}

func Partition(nums []int, l, r int) int {
	v := nums[(l+r)/2]
	i, j := l, r
	for i <= j {
		for nums[i] < v {
			i++
		}
		for nums[j] > v {
			j--
		}

		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return j
}

func main() {
	nums := []int{1, -1, 2, 7, 5, 9}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println("nums = ", nums)
}
