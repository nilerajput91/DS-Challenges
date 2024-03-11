/* Given a string S. The task is to count the number of substrings which contains equal number of lowercase and uppercase letters.

Example 1:

Input:
S = "gEEk"
Output: 3
Explanation: There are 3 substrings of
the given string which satisfy the
condition. They are "gE", "gEEk" and "Ek".
Example 2:

Input:
S = "WomensDAY"
Output: 4
Explanation: There are 4 substrings of
the given string which satisfy the
condition. They are "Wo", "ensDAY",
"nsDA" and "sD"
Your Task:
The task is to complete the function countSubstring() which takes the string S as input parameter and returns the number of substrings which contains equal number of uppercase and lowercase letters.

Expected Time Complexity: O(N*N)
Expected Auxiliary Space: O(1)

Constraints:
1 ≤ |S| ≤ 103

*/

package main

import "fmt"

func main() {
	S := "gEEk"

	fmt.Printf("count of Substring %d \n:", countSubString(S))

}

func countSubString(s string) int {
	count := 0

	n := len(s)

	for i := 0; i < n; i++ {
		lowerCount, upperCount := 0, 0

		for j := i; j < n; j++ {

			if s[j] >= 'a' && s[j] <= 'z' {
				lowerCount++

			} else {
				upperCount++
			}
			if lowerCount == upperCount {
				count++
			}
		}
	}

	return count

}
