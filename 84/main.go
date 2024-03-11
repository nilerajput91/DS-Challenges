package main

import "fmt"

func main() {
	N := 10
	num := []int{6,1,2,8,3,4,7,10,5}
	fmt.Printf("missing number is: %d \n",missingNumber(num,N))

}

func missingNumber(arr []int, num int) int {
	expectedSum := num * (num + 1) / 2

	actualSum := 0

	for _, num1 := range arr {
		actualSum += num1

	}

	return expectedSum - actualSum

}
