/* Given two strings text1 and text2, return the
length of their longest common subsequence.
*/

package main

import "fmt"

func main() {

	str1 := "ace"
	str2 := "abcde"

	fmt.Printf("longest substring using :%v\n", longestCommonSubsequence(str1, str2))

}
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)

	//create 2d array to store the len of longest common substring
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

		}
	}

	return dp[m][n]
}
