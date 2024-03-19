/*
The Substring with Concatenation of All Words problem is a challenging problem where
you are given a string s and a list of words words, and you need to
find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once
and without any intervening characters
*/
package main

import "fmt"

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println("Input:", s, words)

	fmt.Printf("Output:%v\n", findSubstrings(s, words))

}

func findSubstrings(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	// Calculate the length of each word and the total length of all words
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	var result []int

	// Iterate through each starting index i up to len(s) - totalLen
	for i := 0; i < len(s)-totalLen; i++ {

		seen := make(map[string]int)

		j := 0

		for j < len(words) {
			// Extract the current word from the string s
			word := s[i+j*wordLen : i+(j+1)*wordLen]

			// Check if the current word is in the wordCount map
			if count, ok := wordCount[word]; ok {
				// Increment the count of the current word in the seen map
				seen[word]++

				// If we've seen the current word more times than its count, break
				if seen[word] > count {
					break
				}
			} else {
				// If the current word is not in the wordCount map, break
				break
			}
			// Move to the next word in the words list
			j++

		}
		// If we've successfully iterated through all words in the words list
		if j == len(words) {
			result = append(result, i)
		}
	}
	return result

}
