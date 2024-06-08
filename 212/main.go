//To find a peak element in an array using the binary search approach

package main

import "fmt"

func main() {
	ele := []int{1, 2, 3, 1}
	peakElement := findPeakElement(ele)

	fmt.Printf("peak element %v", peakElement)

}

func findPeakElement(nums []int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			// If the element before mid is smaller, mid is a peak
			// Otherwise, move to the left side of mid
			right = mid
		} else {
			// If the element after mid is smaller, mid is a peak
			// Otherwise, move to the right side of mid
			left = mid + 1
		}
	}
	return left
}
