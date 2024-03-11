package main

import "fmt"

func main() {

	S := "aaaabaa"
	result := longestpaildrome(S)

	fmt.Println(result)
}

func longestpaildrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {

		dp[i] = make([]bool, n)
	}

	// All substrings of length 1 are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	start, maxLenght := 0, 1

	for i := 0; i < n-1; i++ {

		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLenght = 2
		}
	}

	for length := 3; length <= n; length++ {

		for i := 0; i <= n-length; i++ {

			j := i + length - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLenght = length
			}
		}
	}

	return s[start : start+maxLenght]

}
