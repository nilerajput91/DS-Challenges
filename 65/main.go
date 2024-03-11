package main

import (
	"fmt"
	"sort"
)

func main() {

	inputString := "STRING"
	strBytes := []byte(inputString)

	sort.Slice(strBytes, func(i, j int) bool {
		return strBytes[i] < strBytes[j]
	})

	sortedString := string(strBytes)

	Rank := laxigraphicRank(inputString)
	fmt.Printf("laxigrphacially sorted string is %s", sortedString)

	fmt.Printf("\n laxigraphically input string : %s \n Rank final string :%d \n", inputString, Rank)

}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func laxigraphicRank(str string) int {
	length := len(str)

	fact := factorial(length)

	Rank := 1

	for i := 0; i < length; i++ {
		smallerRight := 0

		for j := i + 1; j < length; j++ {
			if str[j] < str[i] {
				smallerRight++
			}
		}
		Rank += smallerRight * fact / (length - i)
	}

	return Rank

}
