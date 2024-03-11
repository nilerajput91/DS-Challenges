package main

import "fmt"

func main() {

	a, b := 10, 15

	result := gcd(a, b)

	fmt.Println("Gcd of given numbers:", result)

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
