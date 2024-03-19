//Anagram

package main

import "fmt"

func main() {
	s1 := "anagram"
    t1 := "nagaram"
    fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", s1, t1)
    fmt.Println("Output:", isAnagram(s1, t1))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charFreq := make(map[rune]int)

	for _, char := range s {
		charFreq[char]++
	}

	for _, char := range t {

		if freq, ok := charFreq[char]; !ok || freq == 0 {
			return false
		}
		charFreq[char]--

	}
	return true

}
