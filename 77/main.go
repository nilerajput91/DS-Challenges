package main

import "fmt"

func main() {
	n:=8

	num:=countTrailingZero(n)
	fmt.Println("number:",num)

}

func countTrailingZero(n int) int {
	res := 0

	for i := 5; i < n; i = i * 5 {
		res = res +n / i
	}

	return res

}
