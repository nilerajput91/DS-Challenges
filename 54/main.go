package main

import "fmt"

func main() {

	s1 := "ab"
	s2 := "abcd"

	result := isSubsequence(s1, s2)

	if result {
		fmt.Println("its subsequence ", result)
	} else {
		fmt.Println("its not susequence", result)
	}

}

// isSubsequence checks if s1 is a subsequence of s2.
func isSubsequence(s1, s2 string) bool {
	// Initialize indices for s1 and s2.
	sIndex := 0
	tIndex := 0

	// Iterate over s2 until we reach the end of s1 or s2.
	for tIndex < len(s2) && sIndex < len(s1) {
		// If the characters at the current indices match, move to the next character in s1.
		if s1[sIndex] == s2[tIndex] {
			sIndex++
		}
		// Move to the next character in s2.
		tIndex++
	}

	// If we've reached the end of s1, it is a subsequence of s2.
	return sIndex == len(s1)
}
