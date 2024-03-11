package main

import "fmt"

func main() {

	input := "cccccdddllla"

	resultMap := charFrequency(input)

	for char, count := range resultMap {
		fmt.Printf("charcter:%c : freq of char:%d \n", char, count)
	}

}

// charFrequency returns a map containing the frequency of each character in the input string.
func charFrequency(input string) map[rune]int {
	// Initialize a map to store the frequency of each character.
	charFreq := make(map[rune]int)

	// Iterate over each character in the input string.
	for _, char := range input {
		// Increment the frequency count for the current character.
		charFreq[char]++
	}

	// Return the map containing the frequency of each character.
	return charFreq
}
