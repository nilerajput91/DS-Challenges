package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaa"
	fmt.Printf("pailndrome %v\n", isPalidrome(str))

}

func isPalidrome(str string) bool {
	str = strings.ToLower(str)
	isAplhaNumeric := func(char byte) bool {
		return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
	}
	i, j := 0, len(str)-1
	for i < j {
		for i < j && !isAplhaNumeric(str[i]) {
			i++

		}

		for i < j && !isAplhaNumeric(str[j]) {
			j--
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--

	}
	return true

}
