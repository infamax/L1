package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

var incorrectPosition = errors.New("incorrect position byte")

func ChangeByteNum(num int64, i int, val int) (int64, error) {
	if lengthNum(num) <= i {
		return 0, incorrectPosition
	}

	var mask int64 = 1 << i

	if val == 0 {
		mask = ^mask
		return num & mask, nil
	}

	return num | mask, nil
}

func lengthNum(num int64) int {
	return len(strconv.FormatInt(num, 2))
}

func main() {
	var test int64 = 10
	fmt.Println(ChangeByteNum(test, 1, 0))
	fmt.Println(ChangeByteNum(test, 2, 1))
	fmt.Println(ChangeByteNum(test, 3, 1))
}
