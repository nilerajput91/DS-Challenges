/* Given a string with brackets ('[' and ']') and the index of an opening bracket. Find the index of the corresponding closing bracket.

Example 1:

Input:
S = "[ABC[23]][89]"
pos = 0
Output: 8
Explanation: [ABC[23]][89]
The closing bracket corresponding to the
opening bracket at index 0 is at index 8.

Example 2:

Input:
S = "ABC[58]"
pos = 3
Output: 6
Explanation: ABC[58]
The closing bracket corresponding to the
opening bracket at index 3 is at index 6.

Your Task:
You don't need to read input or print anything. Your task is to complete the function closing() which takes a string S and a position pos as inputs and returns the index of the closing bracket corresponding to the opening bracket at index pos.


Expected Time Complexity: O(|S|).
Expected Auxiliary Space: O(1).


Constraints:
1 <= N <= 2*105 */

package main

import "fmt"

func main() {
	S := "[ABC[23]][89]"
	pos := 0

	fmt.Printf("index of %d\n", closing(S, pos))

}

func closing(s string, pos int) int {

	stack := []int{}
	for i := 0; i < len(s); i++ {

		if s[i] == '[' {
			stack = append(stack, i)
		} else if s[i] == ']' {
			if len(stack) == 0 {
				return -1
			}

			if stack[len(stack)-1] == pos {
				return i
			}
			stack = stack[:len(stack)-1]

		}

	}
	return -1

}
