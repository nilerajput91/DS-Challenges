package main

import "fmt"

func main() {
	input := "abcabcbb"
	result := logestSubestring(input)
	fmt.Printf("Length of the longest substring with distinct characters: %d\n", result)
}

func logestSubestring(s string) int {

	charindex := make(map[rune]int)
	maxLen := 0
	start := 0

	for end, char := range s {

		if index, found := charindex[char]; found {
			start = max(start, index+1)
		}

		//update char index for end

		charindex[char] = end

		maxLen = max(maxLen, end-start+1)

	}

	return maxLen

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
