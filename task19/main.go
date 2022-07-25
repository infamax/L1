package main

import "fmt"

/*

Разработать программу, которая
переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode
*/

func ReverseString(str string) string {
	strRune := []rune(str)
	for i := 0; i < len(strRune)/2; i++ {
		strRune[i], strRune[len(strRune)-i-1] = strRune[len(strRune)-i-1], strRune[i]
	}
	return string(strRune)
}

func main() {
	str := "главрыба"
	fmt.Println(ReverseString(str))
}
