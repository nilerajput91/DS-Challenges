//shortest palidrome

package main

import "fmt"

func main() {
	// Example usage
	input := "aacecaaa"
	fmt.Printf("Shortest Palindrome of %s: %s\n", input, shortestPalindrome(input))

}

// shortestPalindrome finds the shortest palindrome by adding characters in front of the given string.
func shortestPalindrome(str string) string {
	// Find the longest palindrome substring starting from index 0
	j := len(str) - 1

	for j >= 0 && !isPalidrome(str[:j+1]) {
		j--
	}
	// Build the shortest palindrome by adding characters from index j+1 to the beginning of the string
	prefix := reverse(str[j+1:])
	return prefix + str

}
func isPalidrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// reverse string
func reverse(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}
