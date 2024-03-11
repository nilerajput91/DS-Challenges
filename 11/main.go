package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str:="xxxxyyzz"
	fmt.Println(checkBalance(str))

}

func checkBalance(s string) bool {

	charFrequency := make(map[rune]int)
	maxFrequency := 0

	for _, char := range s {

		charFrequency[char]++

		if charFrequency[char] > maxFrequency {
			maxFrequency = charFrequency[char]
		}
	}

	if maxFrequency*(maxFrequency-1)/2 == utf8.RuneCountInString(s)-maxFrequency {
		return true
	}

	for _, freq := range charFrequency {
		if freq == 1 {
			return true
		}
	}

	return false
}
