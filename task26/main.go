package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func IsUniqueString(str string) bool {
	m := make(map[rune]struct{})
	str = strings.ToLower(str)
	for _, s := range str {
		_, ok := m[s]
		if ok {
			return false
		}
		m[s] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(IsUniqueString("abcd"))
	fmt.Println(IsUniqueString("abCdefAaf"))
	fmt.Println(IsUniqueString("aabcd"))
}
