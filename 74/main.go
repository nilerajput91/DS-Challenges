package main

import "fmt"

func main() {

	result := maxPieces(23, 11, 9, 12)

	fmt.Println("result", result)

}

func maxPieces(n, a, b, c int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}

	result := max(maxPieces(n-a, a, b, c), maxPieces(n-b, a, b, c), maxPieces(n-c, a, b, c))
	if result == -1 {
		return -1
	} else {

		return result + 1
	}

}
