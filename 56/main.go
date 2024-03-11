package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str1 := "listen"
	str2 := "slient"

	fmt.Printf("are Anagram :%v\n", areAnagram(str1, str2))

}

// areAnagram checks if str1 and str2 are anagrams of each other.
func areAnagram(str1, str2 string) bool {
	// Remove spaces and convert both strings to lowercase.
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")

	// If the lengths of the two strings are not equal, they cannot be anagrams.
	if len(str1) != len(str2) {
		return false
	}

	// Convert strings to runes for sorting.
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// Sort the rune slices.
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	// Compare the sorted rune slices converted back to strings.
	return string(runes1) == string(runes2)
}
