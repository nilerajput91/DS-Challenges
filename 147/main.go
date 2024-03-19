/* Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example:
go
Copy code
Input: s = "egg", t = "add"
Output: true
Explanation: "e" can be replaced with "a" and "g" can be replaced with "d".

Input: s = "foo", t = "bar"
Output: false
Explanation: No possible mapping for characters. */

package main

import "fmt"

func main() {

	s1 := "egg"
	s2 := "add"

	fmt.Printf("isIsomoprhics:%v\n", isIsomorphics(s1, s2))

}

func isIsomorphics(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]

		// Check if the mapping is already established
		if mappedSChar, ok := sMap[tChar]; ok && mappedSChar != sChar {
			return false
		}

		if mappedTChar, ok := tMap[sChar]; ok && mappedTChar != tChar {
			return false
		}
		// Establish the mapping
		sMap[tChar] = sChar
		tMap[sChar] = tChar

	}

	return true

}
