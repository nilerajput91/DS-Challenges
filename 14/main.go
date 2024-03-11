package main

import "fmt"

func main() {

	input:=[]int{1,3,3}
	fmt.Println(missingNumber(input))

}

func missingNumber(arr []int) (int, int) {

	n := len(arr)

	sum := 0
	sumOfSquare := 0

	for i := 0; i < n; i++ {

		sum += arr[i]
		sumOfSquare += arr[i] * arr[i]

	}

	expectedSum := n * (n + 1) / 2

	expectedSumOfSquare := n * (n + 1) * (2*n + 1) / 6

	diffnSum := expectedSum - sum
	difffnSumOfSquares := expectedSumOfSquare - sumOfSquare

	A := (diffnSum + difffnSumOfSquares/diffnSum) / 2
	B := A - diffnSum
	return B, A

}
