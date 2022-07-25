package main

import (
	"fmt"
	"reflect"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10)
	и выведет их квадраты в stdout.
*/

func SquareNumsV1(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(nums))
		for i := 0; i < len(nums); i++ {
			go func(i int) {
				defer wg.Done()
				ch <- nums[i] * nums[i]
			}(i)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	nums := []int{2, 4, 6, 8}

	for val := range SquareNumsV1(nums) {
		fmt.Println("val = ", val)
	}

	fmt.Println(reflect.DeepEqual([]int{2, 4, 6, 8}, []int{8, 6, 2, 4}))
}
