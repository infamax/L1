package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func IntervalsTemperature(temps []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, temp := range temps {
		m[(int(math.Trunc(temp))/10)*10] = append(m[(int(math.Trunc(temp))/10)*10], temp)
	}
	return m
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	intervalsTemps := IntervalsTemperature(temps)
	fmt.Println(intervalsTemps)
}
