/* Given a string S of opening and closing brackets '(' and ')' only. The task is to find an equal point. An equal point is an index such that the number of closing brackets on right from that point must be equal to number of opening brackets before that point.


Example 1:

Input: str = "(())))("
Output: 4
Explanation:
After index 4, string splits into (())
and ))(. Number of opening brackets in the
first part is equal to number of closing
brackets in the second part.
Example 2:
Input : str = "))"
Output: 2
Explanation:
As after 2nd position i.e. )) and "empty"
string will be split into these two parts:
So, in this number of opening brackets i.e.
0 in the first part is equal to number of
closing brackets in the second part i.e.
also 0.

Your Task:
You don't need to read input or print anything. Your task is to complete the function findIndex() which takes the string S as inputs and returns the occurrence of the equal point in the string.


Expected Time Complexity: O(N)
Expected Auxiliary Space: O(N) */

package main

import "fmt"

func main() {
	str := "(())))("

	fmt.Printf("index of closing bracket:%d\n", findIndex(str))

}

func findIndex(s string) int {

	openingCount := 0
	totalClosingCount := 0

	// Count total number of closing brackets
	for _, ch := range s {
		if ch == ')' {
			totalClosingCount++
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openingCount++

		} else {
			totalClosingCount--
		}

		if totalClosingCount == openingCount {
			return i
		}
	}
	return -1

}
