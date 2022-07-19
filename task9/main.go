package main

import (
	"fmt"
	"sync"
)

/*
	Разработать конвейер чисел. Даны два канала:
	в первый пишутся числа (x) из массива,
	во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func GetNumsChan(nums []int) <-chan int {
	numsCh := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(nums))
		for i := 0; i < len(nums); i++ {
			go func(i int) {
				defer wg.Done()
				numsCh <- nums[i]
			}(i)
		}
		wg.Wait()
		close(numsCh)
	}()
	return numsCh
}

func X2Nums(numsCh <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for val := range numsCh {
			res <- val * 2
		}
		close(res)
	}()
	return res
}

func main() {
	nums := []int{2, 4, 6, 8}
	for val := range X2Nums(GetNumsChan(nums)) {
		fmt.Println("val = ", val)
	}
}
