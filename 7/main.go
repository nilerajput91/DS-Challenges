package main

import "fmt"

func main() {
	fmt.Println(isSequenece("190"))
}

func isSequenece(str string) bool {

	for _, char := range str {

		if char != '0' && char != '1' {
			return false
		}
	}

	return true
}
