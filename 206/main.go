package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 3, 2, 4}
	fmt.Println(minDiffernece(nums)) // Output: 0

}

func minDiffernece(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	if n <= 3 {
		return 0
	}
	// Compare four possible scenarios:
	// 1. Remove the three largest elements
	// 2. Remove the three smallest elements
	// 3. Remove the two smallest and one largest elements
	// 4. Remove the one smallest and two largest elements
	return min(
		nums[n-1]-nums[3],
		nums[n-2]-nums[2],
		nums[n-3]-nums[1],
		nums[n-4]-nums[0],
	)
}

func min(a, b, c, d int) int {
	return minOfTwo(minOfTwo(a, b), minOfTwo(c, d))
}
func minOfTwo(a, b int) int {
	if a < b {
		return a
	}

	return b
}
