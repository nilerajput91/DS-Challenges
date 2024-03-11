package main

import "fmt"

func main() {

	element := []int{3, 10, 20, 30}
	secLargestEle := secondLargest(element)
	fmt.Println(secLargestEle)

}

// secondLargest finds the index of the second largest element in arr.
// If arr has less than 2 elements, it returns -1.
func secondLargest(arr []int) int {
	// Check if arr has less than 2 elements, in which case return -1.
	if len(arr) < 2 {
		return -1
	}

	// Initialize variables to track the indices of the largest and second largest elements.
	secondLargest := -1
	largest := 0

	// Iterate through the array starting from the second element.
	for i := 1; i < len(arr); i++ {
		// If the current element is greater than the largest element,
		// update the secondLargest index to the current largest index
		// and update the largest index to the current index.
		if arr[i] > arr[largest] {
			secondLargest = largest
			largest = i
		} else if arr[i] != arr[largest] {
			// If the current element is not equal to the largest element,
			// and the secondLargest index is -1 or the current element is greater than the second largest element,
			// update the secondLargest index to the current index.
			if secondLargest == -1 || arr[i] > arr[secondLargest] {
				secondLargest = i
			}
		}
	}

	// Return the index of the second largest element.
	return secondLargest
}
