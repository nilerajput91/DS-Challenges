// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {

	// Example 1
	s1 := "geeksforgeeks"
	s2 := "forgeeksgeeks"
	fmt.Println(isRotated(s1, s2)) // Output: true

}

func isRotated(s1, s2 string) bool {

	if len(s1) != len(s2) || len(s1) == 0 {

		return false
	}

	concat := s1 + s2
	return strings.Contains(concat, s1)

}
