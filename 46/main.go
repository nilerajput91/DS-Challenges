package main

import "fmt"

func main() {

	arr := []int{1, 4, 20, 3, 10, 5}
	targetSum := 5

	subArray := sumSubarray(arr, targetSum)

	if len(subArray) > 0 {
		fmt.Printf("longest subarray of sum %d:%v=", targetSum, subArray)
	} else {
		fmt.Println("longest sub array")
	}

}

// sumSubarray finds a subarray of arr that sums to targetSum using the "Sliding Window" technique.
func sumSubarray(arr []int, targetSum int) []int {
	// Initialize variables for the start and end of the window, and the current sum.
	start := 0
	currentSum := 0
	end := 0

	// Loop until the end of the array is reached.
	for end < len(arr) {
		// Add the element at index end to the current sum.
		currentSum += arr[end]

		// If the current sum is greater than the target sum, move the start of the window to the right
		// and subtract the element at the previous start position from the current sum, until the current sum
		// is less than or equal to the target sum.
		for currentSum > targetSum && start <= end {
			currentSum -= arr[start]
			start++
		}

		// If the current sum is equal to the target sum, return the subarray from start to end,
		// as it is the subarray that sums to the target sum.
		if currentSum == targetSum {
			return arr[start : end+1]
		}

		// Move the end of the window to the right to expand the window.
		end++
	}

	// If no subarray is found, return an empty array.
	return []int{}
}
