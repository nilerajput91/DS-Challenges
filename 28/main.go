//left most repeating char

package main

import "fmt"

func main() {

	input := "abacabad"
	character, index := leftMostRepatingChar(input)

	if index {
		fmt.Printf("Leftmost repeating character: %c at index %d\n", character, index)
	} else {
		fmt.Println("No repeating characters found.")
	}

}

func leftMostRepatingChar(s string) (rune, bool) {


	charIndex := make(map[rune]int)

	for i, char := range s {
		if _, found := charIndex[char]; found {
			return char, found
		}
		charIndex[char] = i

	}

	return 0, false

}
