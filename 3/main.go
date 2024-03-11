// longest valid parnethieis
// Click here and start typing.
package main

import "fmt"

func main() {
	s1 := "("
	fmt.Println(longestValidParthensis(s1))
}

func longestValidParthensis(str string) int {
	stack := []int{-1}
	maxLength := 0
	for i := 0; i < len(str); i++ {

		if str[i] == '(' {

			stack = append(stack, i)
		} else {

			stack = stack[:len(stack)-1]

			if len(stack) == 0 {

				stack = append(stack, i)
			} else {

				currentlength := i - stack[len(stack)-1]
				maxLength = max(maxLength, currentlength)
			}
		}
	}

	return maxLength

}

func max(a, b int) int {

	if a > b {
		return a
	} else {

		return b
	}
}
