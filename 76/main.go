package main

import "fmt"

func main() {

	n := 333

	isPali := isPalidrome(n)

	if isPali {
		fmt.Println("palidrome")
	} else {
		fmt.Println("not palidrome")
	}

}

func isPalidrome(n int) bool {
	rev := 0
	temp := n

	for temp != 0 {
		rem := temp % 10
		rev = rev*10 + rem
		temp = temp / 10
	}

	return rev == n
}
