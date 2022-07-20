package main

import (
	"fmt"
	"math"
)

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
