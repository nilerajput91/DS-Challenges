package main

import "fmt"

func main() {

	arr := []int{4, 2, -3, 1, 6}

	isZero := subArrayOfZeros(arr)
	if isZero {
		fmt.Println("given array sub array zero")
	} else {
		fmt.Println("given array sub array have no zero")

	}

}

func subArrayOfZeros(arr []int) bool {
	sumMap := make(map[int]bool)

	sum := 0

	for _, num := range arr {

		sum += num

		if sum == 0 || sumMap[sum] {
			return true
		}

		sumMap[sum] = true

	}

	return false
}
