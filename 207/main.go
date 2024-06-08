// The problem of finding buddy strings involves comparing two strings to determine if they can become equal by swapping exactly two characters
package main

import "fmt"

func main() {
	A := "ab"
	B := "ba"
	fmt.Println(buddyString(A, B)) // Output: true

}

func buddyString(A, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		seen := make(map[byte]bool)

		for i := 0; i < len(A); i++ {
			// Found duplicate characters, swapping them will result in equal strings
			if seen[A[i]] {
				return true
			}

			seen[A[i]] = true
		}
		// No duplicate characters, cannot swap any characters to make them equal
		return false

	}
	// Find the indices where A and B differ
	diffIndices := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diffIndices = append(diffIndices, i)
		}
	}
	// Check if there are exactly two differing indices
	if len(diffIndices) != 2 {
		return false
	}
	// Check if swapping characters at the differing indices results in equal strings

	return A[diffIndices[0]] == B[diffIndices[1]] && A[diffIndices[1]] == B[diffIndices[0]]

}
