//To determine if a string is a subsequence of another string

package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

}

func isSubsequence(s, t string) bool {
	if len(s) == 0 {
		return true
	}

	sIndex := 0
	tIndex := 0

	if sIndex < len(s) {
		if s[sIndex] == t[tIndex] {
			sIndex++

			if sIndex == len(s) {
				return true
			}
		}
		tIndex++
	}
	return true
}
