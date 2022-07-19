package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 1, 2
	// Первый способ
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Второй способ
	x, y := 1, 2
	x -= y
	y += x
	x -= y
	x = -x
	fmt.Printf("x = %d, y = %d\n", x, y)
}
