package main

import "fmt"

func main() {

	s := "{{}[}"

	fmt.Println(minReversal(s))

}

func minReversal(s string) int {

	n := len(s)

	if n%2 != 0 {
		return -1

	}

	openCount, closeCount := 0, 0

	for i := 0; i < n; i++ {

		if s[i] == '{' || s[i] == '[' {

			openCount++
		} else if s[i] == '}' || s[i] == ']' {

			if openCount == 0 {
				closeCount++
			} else {

				openCount++
			}

		}

	}

	reversal := (openCount+1)/2 + (closeCount+1)/2

	return reversal
}
