package main

import "fmt"

func main() {

	text := "ABCABC"
	pattern := "ABC"

	result := NaivepatternSearch(text, pattern)

	if len(result) > 0 {
		fmt.Printf(" \n text :%s match pattern:%s \n", text, pattern)
	} else {
		fmt.Println("pattern is not match")
	}

}

func NaivepatternSearch(text, pattern string) []int {
	var occurrences []int

	n := len(text)
	m := len(pattern)

	i := 0

	for i <= n-m {
		j := 0

		for j < m && pattern[j] == text[i+j] {
			j++
		}

		if j == m {
			occurrences = append(occurrences, i)
			i += m
		} else if j == 0 {
			i++

		} else {
			i += j
		}
	}

	return occurrences
}
