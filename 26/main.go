package main

import (
	"fmt"
	"slices"
)

func main() {
	str := "Hello world"
	result := reverse(str)
	fmt.Println(result)

	sl := []string{"hello"}
	slices.Reverse(sl)
	fmt.Println(sl)

}

func reverse(str string) string {

	strRune := []rune(str)

	lentght := len(strRune)

	for i, j := 0, lentght-1; i < j; i, j = i+1, j-1 {

		strRune[i], strRune[j] = strRune[j], strRune[i]

	}

	return string(strRune)

}
