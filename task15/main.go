package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может
привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.
	var justString string - использование глобальной переменной
	func someFunc() {
		v := createHugeString(1 << 10) - выделяется больший блок памяти, который не используется, можно сразу выделить память
		под весь блок
		justString = v[:100]
	}
	func main() {
		someFunc()
	}
*/

func createHugeString(length int) string {
	var res strings.Builder
	for i := 0; i < length; i++ {
		res.WriteString("a")
	}
	return res.String()
}

func someFunc() string {
	return createHugeString(100)
}

func main() {
	fmt.Println(someFunc())
}
