package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел:
	2,4,6,8,10. Найти сум
	му их квадратов(2^2+3^2+4^2….) с
	использованием конкурентных вычислений.

*/

func SumSquareNums(nums []int) int {
	squareNumsCh := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(nums))
		for i := 0; i < len(nums); i++ {
			go func(i int) {
				defer wg.Done()
				squareNumsCh <- nums[i] * nums[i]
			}(i)
		}
		wg.Wait()
		close(squareNumsCh)
	}()

	res := 0

	for val := range squareNumsCh {
		res += val
	}

	return res
}

func main() {
	nums := []int{2, 4, 6, 8}

	fmt.Println("SumSquareNums = ", SumSquareNums(nums))
}
