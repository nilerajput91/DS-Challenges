// Find First and Last Position of Element in Sorted Array
package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := searchRange(nums, target)
	fmt.Println("First and last positions of", target, "in the sorted array:", result)

}

func findFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1

			if nums[mid] == target {
				result = mid
			}
		} else {
			left = mid + 1
		}
	}
	return result
}

func findLastPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1

			if nums[mid] == target {
				result = mid
			}

		} else {
			right = mid - 1
		}
	}
	return result
}
func searchRange(nums []int, target int) []int {
	first := findFirstPosition(nums, target)
	last := findLastPosition(nums, target)
	return []int{first, last}

}
