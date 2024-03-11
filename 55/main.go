package main

import "fmt"

func main() {

	teststring := "aaa"

	isTrue := isPalindrome(teststring)

	if isTrue {
		fmt.Println("the above string is palidrome")
	} else {
		fmt.Println("its not palidrome")
	}

}

// isPalindrome checks if the input string s is a palindrome.
func isPalindrome(s string) bool {
	// Calculate the length of the input string.
	length := len(s)

	// Iterate over the first half of the string.
	for i := 0; i < length/2; i++ {
		// Compare the character at index i with its corresponding character from the end of the string.
		// If they are not equal, the string is not a palindrome.
		if s[i] != s[length-1-i] {
			return false
		}
	}

	// If all characters match, the string is a palindrome.
	return true
}
