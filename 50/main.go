package main

import "fmt"

func main() {

	arr := []int{12, 12}
	removeDuplicateArray := removeDuplicate(arr)
	fmt.Println(removeDuplicateArray)

}

// removeDuplicate removes duplicate elements from the array arr in-place and returns the length of the resulting array.
func removeDuplicate(arr []int) int {
	// Initialize a variable to track the position where the next non-duplicate element should be placed.
	result := 1

	// Iterate through the array starting from the second element.
	for i := 1; i < len(arr); i++ {
		// If the current element is different from the previous non-duplicate element,
		// copy the current element to the position indicated by result and increment result.
		if arr[i] != arr[result-1] {
			arr[result] = arr[i]
			result++
		}
	}

	// Return the length of the resulting array, which is equal to result.
	return result
}
