package main

import "fmt"

func main() {

	n:=4

	result:=getSum(n)
	fmt.Println(result)

}

// recursive
func getSum(n int) int {
	if n == 0 {
		return 0
	}

	return n + getSum(n-1)
}
