//To find the index of the first occurrence of a substring in a string

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	substring := "World"

	Index := strings.Index(str, substring)

	if Index != -1 {
		fmt.Printf("found at %d index substring is %s \n", Index, str)
	} else {
		fmt.Printf("not found at %d index substring is %s\n ", Index, str)

	}

}
