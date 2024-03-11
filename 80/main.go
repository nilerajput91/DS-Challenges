package main

import "fmt"

func main() {
	isPrime := isPrime(121)
	fmt.Printf("The above number is Prime: %v", isPrime)

}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i = i + 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true

}
