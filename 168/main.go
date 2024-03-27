/* Problem Statement:
Given a balanced parentheses string S, compute the score of the string based on the following rules:

() has a score of 1.
AB has a score of A + B, where A and B are balanced parentheses strings.
(A) has a score of 2 * A, where A is a balanced parentheses string.
The score of an empty string is 0.

Example 1:
Input: "()"
Output: 1

Example 2:
Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
Example 4:

Input: "(()(()))"
Output: 6
*/

package main

import "fmt"

func main() {
	Input := "()()"
	fmt.Printf("result:%v\n", countParathesis(Input))

}

func countParathesis(s string) int {
	stack := []int{0}

	for _, str := range s {
		if str == '(' {
			stack = append(stack, 0) // Push 0 for the '('
		} else {
			top := stack[len(stack)-1]    // Get the top of the stack
			stack := stack[:len(stack)-1] //Pop the stack

			if top == 0 {
				stack[len(stack)-1]++ // Update the previous score
			} else {
				stack[len(stack)-1] += 2 * top // Update the previous score
			}
		}
	}
	score := 0

	for _, s := range stack {
		score += s
	}
	return score
}
