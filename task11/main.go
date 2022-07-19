package main

import (
	"fmt"
	"sort"
)

func Intersection(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	prev := -1
	f, s := 0, 0
	for i := 0; i < len(nums1)+len(nums2); i++ {
		if f >= len(nums1) {
			break
		}

		if s >= len(nums2) {
			break
		}

		if nums1[f] < nums2[s] {
			f++
		} else if nums1[f] > nums2[s] {
			s++
		} else {
			if nums1[f] != prev {
				res = append(res, nums1[f])
				prev = nums1[f]
			}
			s++
			f++
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(Intersection(nums1, nums2))
}
