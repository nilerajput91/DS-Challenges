package main

import (
	"fmt"
	"math"
)

const prime = 100

func main() {

	text := "ABCCDDAEFGCABCCDDEFG"
	pattern := "CCD"

	occurrences := rabinKarapSearch(text, pattern)
	if len(occurrences) > 0 {
		fmt.Printf("pattern found at position:%v\n", occurrences)
	} else {
		fmt.Println("pattern not found in the text")
	}
}

func rabinKarapSearch(text, pattern string) []int {
	var occurrences []int

	n := len(text)
	m := len(pattern)

	patternHash := hash(pattern, m)
	textHash := hash(text[:m], m)

	for i := 0; i < n-m; i++ {
		if patternHash == textHash && text[i:i+m] == pattern {
			occurrences = append(occurrences, i)
		}

		if i+m < n {
			textHash = recalculateHash(textHash, int(text[i]), int(text[i+m]), m)
		}
	}

	return occurrences
}

func hash(s string, len int) int {
	result := 0
	for i := 0; i < len; i++ {
		result += int(s[i]) * int(math.Pow(prime, float64(i)))
	}

	return result

}

func recalculateHash(oldHash, oldChar, newChar, len int) int {
	return (oldHash-oldChar)/prime + newChar*int(math.Pow(prime, float64(len-1)))
}
