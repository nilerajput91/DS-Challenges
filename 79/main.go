package main

import "fmt"

func main() {
	resultGcd := gcd(2, 4)
	fmt.Println(resultGcd)

	lcmResult := lcm(2, 4)
	fmt.Println(lcmResult)

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// lcm least common divsior
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
