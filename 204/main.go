// first occurances of string
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	substr := "world"
	index := firstOccurances(s, substr)
	if index != -1 {
		fmt.Printf("First occurrence of '%s' in '%s' is at index %d\n", substr, s, index)
	} else {
		fmt.Printf("'%s' is not found in '%s'\n", substr, s)
	}

}

func firstOccurances(s, subStr string) int {
	index := strings.Index(s, subStr)
	return index
}
