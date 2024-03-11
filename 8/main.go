
package main

import "fmt"

func main() {
	fmt.Println(string(firstRepeatedCharacter("strst")))
}

func firstRepeatedCharacter(str string) rune {
	charCount := make(map[rune]int)

	for _, rune := range str {

		if count, exist := charCount[rune]; exist {

			fmt.Println(count)

			return rune

		}

		charCount[rune] = 1
	}

	return -1
}
