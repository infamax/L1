package main

import (
	"fmt"
	"strings"
)

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
