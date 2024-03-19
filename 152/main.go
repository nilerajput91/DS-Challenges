package main

import (
	"fmt"
	"strings"
)

func main() {
	str:="HelloWorld"

	fmt.Printf("lenOfLastWord:%v\n",lengthOfLastWord(str))

}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	lastSpace := strings.LastIndex(s, " ")

	if lastSpace == -1 {
		return len(s)
	}

	return len(s) - lastSpace - 1

}
