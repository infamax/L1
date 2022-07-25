package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func ReverseWordsInStrings(s string) string {
	if len(s) == 0 {
		return s
	}

	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := " snow dog sun"
	fmt.Println(ReverseWordsInStrings(s))
}
