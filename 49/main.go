package main

import "fmt"

func main() {

	input := []int{10, 30, 2, 3}

	revArray := reverseArray(input)
	fmt.Println()
	fmt.Println("final output and reverse array is :", revArray)

}

// reverseArray reverses the elements of arr in place and returns the reversed array.
func reverseArray(arr []int) []int {
	// Initialize variables to track the indices of the first and last elements.
	low := 0
	high := len(arr) - 1
	// Temporary variable to hold the value during the swap.
	temp := 0

	// Loop until the low index is less than the high index.
	for low < high {
		// Swap the elements at the low and high indices.
		temp = arr[low]
		arr[low] = arr[high]
		arr[high] = temp
		// Move the low index to the right and the high index to the left.
		low++
		high--
	}

	// Return the reversed array.
	return arr
}
