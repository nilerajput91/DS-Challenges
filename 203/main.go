package main

import "fmt"

func main() {


	fmt.Println(isSubsequence("abc","abge"))

}
func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == s[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
