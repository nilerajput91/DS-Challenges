package main

import "fmt"

func main() {

	arr := []int{1, 0, 0, 0, 12}
	moveZerosToEnd(arr)

	fmt.Println(arr)

}

// moveZerosToEnd moves all zeros in the array arr to the end while maintaining the relative order of non-zero elements.
func moveZerosToEnd(arr []int) {
	// Initialize a variable to track the index where the next non-zero element should be placed.
	nonZeroIndex := 0

	// Iterate through the array.
	for i := 0; i < len(arr); i++ {
		// If the current element is non-zero, swap it with the element at nonZeroIndex
		// and increment nonZeroIndex to prepare for the next non-zero element.
		if arr[i] != 0 {
			arr[i], arr[nonZeroIndex] = arr[nonZeroIndex], arr[i]
			nonZeroIndex++
		}
	}
}
