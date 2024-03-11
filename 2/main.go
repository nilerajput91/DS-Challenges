// Here's a simple Go implementation of a function that checks if two strings match where the first string may contain wildcard characters '*' and '?':
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(isMatch("helloABworld", "hello??world??"))
}

func isMatch(str, pattern string) bool {

	i, j := 0, 0

	starIdx, matchIndx := -1, -1

	for i < len(str) {

		if j < len(pattern) && (pattern[j] == '?' || str[i] == pattern[j]) {
			i++
			j++
		} else if j < len(pattern) && pattern[j] == '*' {
			starIdx = j
			matchIndx = i
			j++
		} else if starIdx != -1 {
			j = starIdx + 1
			matchIndx++
			i = matchIndx

		} else {

			return false
		}

	}

	for j < len(pattern) && pattern[j] == '*' {
		j++

	}

	return j == len(pattern)

}
