package main

import "fmt"

func main() {
	arr := []int{1, 4, 20, 3, 10, 5}
	targetSum := 33

	subArray := longestSubArraySum(arr, targetSum)

	if len(subArray) > 0 {
		fmt.Printf("logest subarry is %d : %v", subArray, targetSum)
	} else {
		fmt.Println("not found")
	}
}

// longestSubArraySum finds the longest subarray in arr that sums to targetSum.
func longestSubArraySum(arr []int, targetSum int) []int {
	// Initialize a map to store the sum of elements up to index i as key and index i as value.
	sumMap := make(map[int]int)
	// Initialize variables to store the length of the longest subarray, the starting index of the subarray,
	// and the current sum of elements.
	maxLength := 0
	startIndex := 0
	currentSum := 0

	// Iterate through the array.
	for i, num := range arr {
		// Add the current element to the current sum.
		currentSum += num

		// If the current sum is equal to the target sum, update the maxLength and startIndex
		// to include the entire array up to this point.
		if currentSum == targetSum {
			maxLength = i + 1
			startIndex = 0
		}

		// If the current sum is not found in the sumMap, add it to the map with its index.
		if _, found := sumMap[currentSum]; !found {
			sumMap[currentSum] = i
		}

		// If the complement of the target sum (i.e., currentSum - targetSum) is found in the sumMap,
		// update maxLength and startIndex if the length of the current subarray is longer.
		if startIndex, found := sumMap[currentSum-targetSum]; found {
			if i-startIndex > maxLength {
				maxLength = i - startIndex
				startIndex = startIndex + 1
			}
		}
	}

	// If no subarray is found, return an empty array.
	if startIndex == -1 {
		return []int{}
	}

	// Return the longest subarray found.
	return arr[startIndex : startIndex+maxLength]
}
