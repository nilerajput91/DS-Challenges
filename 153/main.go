/*
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay), or
it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.
*/
package main

import "fmt"

func main() {
	numbers := []int{19, 2, 7, 10, 111, 89}

	for _, num := range numbers {
		fmt.Printf("isHappy(%v) = %t\n", num, isHappy(num))
	}
}

func isHappy(n int) bool {

	seen := make(map[int]bool)
	// Continue the process until we reach 1 (happy number) or enter a cycle

	for n != 1 {
		// If we have already seen the current number, it means we are in a cycle

		if seen[n] {
			return false
		}

		// Mark the current number as seen

		seen[n] = true
		n = getNext(n)
	}
	// If we reach 1, it is a happy number

	return true

}

func getNext(n int) int {
	// Calculate the sum of the squares of the digits of the number
	sum := 0
	for n > 0 {
		digit := n % 10
		sum = digit * digit
		n /= 10
	}
	return sum
}
