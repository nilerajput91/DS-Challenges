package main

import "fmt"

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	occurrences := kmpSearch(text, pattern)

	if len(occurrences) > 0 {
		fmt.Printf("occurences :%v\n", occurrences)
	} else {
		fmt.Printf("occurences: %v\n", occurrences)
	}
}

func computeLpsArray(pattern string) []int {
	m := len(pattern)

	lps := make([]int, m)

	length := 0

	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func kmpSearch(text, pattern string) []int {

	n := len(text)
	m := len(pattern)

	occurrences := []int{}

	lps := computeLpsArray(pattern)

	i, j := 0, 0

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			occurrences = append(occurrences, i-j)
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}

		}
	}
	return occurrences

}
