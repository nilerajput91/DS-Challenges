package main

import (
	"fmt"
	"sort"
)

func main() {

	text := "geeksforgeeks"
	pattern := "frog"

	occurences := anagramSearch(text, pattern)
	if len(occurences) > 0 {

		fmt.Printf("anagram found at position: %v \n", occurences)
	} else {
		fmt.Printf("occurnaces:%v \n", occurences)
	}

}

type sortRunes []rune

func (s sortRunes) len() int {
	return s.len()
}

func (s sortRunes) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) less(i, j int) bool {
	return s[i] < s[j]
}

func areAnagram(str1, str2 string) bool {
	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str1)
	return sortedStr1 == sortedStr2

}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func anagramSearch(text, pattern string) []int {
	var occurences []int

	n := len(text)
	m := len(pattern)

	for i := 0; i <= n-m; i++ {
		if areAnagram(pattern, text[i:i+m]) {
			occurences = append(occurences, i)

		}

	}
	return occurences
}
