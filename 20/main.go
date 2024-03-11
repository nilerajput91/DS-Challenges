package main

import "fmt"

const MAX_CHAR = 256

func main() {
	s1 := "AABBBCBBAC"
	re := findSubstring(s1) // Output: 3
	fmt.Println(re)
}

func findSubstring(str string) int {
	// MAX_CHAR is the maximum number of ASCII characters
	const MAX_CHAR = 256

	// n is the length of the input string
	n := len(str)

	// freq array stores the frequency of characters in the current window
	freq := make([]int, MAX_CHAR)

	// distinctCount is the count of distinct characters in the current window
	distinctCount := 0

	// left and right are pointers to track the current window
	left, right := 0, 0

	// startInd is the starting index of the smallest substring found so far
	startInd := 0

	// minLength is the length of the smallest substring found so far
	minLength := n + 1

	// Iterate through the string using the right pointer
	for right < n {
		// Increment the frequency of the character at str[right] in freq
		if freq[str[right]] == 0 {
			distinctCount++
		}
		freq[str[right]]++

		// Check if distinctCount reaches MAX_CHAR
		for distinctCount == MAX_CHAR {
			// Update minLength if the current window length is smaller
			if right-left+1 < minLength {
				minLength = right - left + 1
				startInd = left
			}

			// Decrement the frequency of the character at str[left] in freq
			freq[str[left]]--
			// If the frequency becomes 0, decrement distinctCount
			if freq[str[left]] == 0 {
				distinctCount--
			}

			// Move the left pointer to the right to contract the window
			left++
		}

		// Move the right pointer to the right to expand the window
		right++
	}

	// If minLength remains unchanged, no valid substring found, return -1
	if minLength == n+1 {
		return -1
	}

	// Otherwise, return minLength
	return minLength
}
