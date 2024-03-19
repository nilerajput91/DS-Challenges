/* Here's an implementation of the Letter Combinations of a Phone Number problem in Go
with comments explaining each part of the code */

package main

import "fmt"

func main() {

	digit := "23"

	fmt.Printf("letter combinations:%v \n", letterCombinations(digit))

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	backtrack(mapping, &result, "", digits, 0)
	return result
}

func backtrack(mapping map[byte]string, result *[]string, current, digits string, index int) {
	if index == len(digits) {

		// All digits have been processed, add the current combination to the result
		*result = append(*result, current)
		return

	}

	digit := digits[index]
	letters := mapping[digit]

	for i := 0; i < len(letters); i++ {

		// Choose a letter from the current digit's letters
		letter := letters[i]

		// Explore the next digit
		backtrack(mapping, result, current+string(letter), digits, index+1)

	}

}
