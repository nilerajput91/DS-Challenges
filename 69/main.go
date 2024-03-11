package main

import "fmt"

func main() {
	str := "ABC"
	printPermutationWihtoutAB(str)

}

func printPermutationWihtoutAB(str string) {
	permutation := []rune{}
	visited := make([]bool, len(str))
	backTrack(str, permutation, visited)
}

func backTrack(str string, permutation []rune, visited []bool) {

	runeStr := []rune(str)

	if len(str) == len(permutation) {
		fmt.Println(string(permutation))
		return

	}

	for i := 0; i < len(runeStr); i++ {

		if !visited[i] && isValid(permutation, runeStr[i]) {

			visited[i] = true
			permutation = append(permutation, runeStr[i])

			backTrack(string(runeStr), permutation, visited)

			visited[i] = false

			if len(permutation) > 0 {
				permutation = permutation[:len(permutation)-1]
			}
		}
	}
}

func isValid(permutation []rune, nextChar rune) bool {

	if len(permutation) > 0 && permutation[len(permutation)-1] == 'A' && nextChar == 'B' {
		return false
	}

	return true

}
