/* Generate Parentheses is a backtracking problem where you need to generate all combinations of well-formed parentheses given a number n,
which represents the number of pairs of parentheses. */

package main

import "fmt"

func main() {
	n := 3
	resultPar := generateParaenthesis(n)

	for _, value := range resultPar {
		fmt.Println(value)
	}

}

func generateParaenthesis(n int) []string {
	var result []string
	backTrack(&result, "", 0, 0, n)
	return result
}

func backTrack(result *[]string, current string, open, close, max int) {
	// Base case: if the length of the current combination is equal to 2*n, add it to the result
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// Recursive case: add an open parenthesis if the number of open parentheses is less than n

	if open < max {
		backTrack(result, current+"(", open+1, close, max)
	}

	// Recursive case: add a closed parenthesis if the number of closed parentheses is less than the number of open parentheses

	if close < open {
		backTrack(result, current+")", open, close+1, max)
	}
}
