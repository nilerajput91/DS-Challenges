/*
The Ransom Note problem is another classic algorithmic problem. The problem statement goes like this:

You have a ransom note that needs to be constructed from a magazine. You are given a string ransomNote representing the ransom note and another string magazine representing the magazine. You need to determine if you can construct the ransom note using the characters from the magazine. You can use each character in the magazine only once.
*/
package main

import "fmt"

func main() {
	rasomeNote := "aaaa"
	magzine := "aaaa"
	fmt.Printf("ransomeNote:%v\n", canConstruct(rasomeNote, magzine))

}

func canConstruct(ransomNote, magzine string) bool {
	magzineMap := make(map[rune]int)

	for _, char := range magzine {
		magzineMap[char]++
	}

	for _, char := range ransomNote {
		if magzineMap[char] <= 0 {
			return false
		}
		magzineMap[char]--
	}

	return true

}
