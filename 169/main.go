/*
Problem Statement:
Given an array nums which consists of non-negative integers and an integer m,
you can split the array into m non-empty continuous subarrays.
Write an algorithm to minimize the largest sum among these m subarrays.

Example 1:

Input: nums = [7,2,5,10,8], m = 2
Output: 18
Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
Example 2:

Input: nums = [1,2,3,4,5], m = 2
Output: 9
Example 3:

Input: nums = [1,4,4], m = 3
Output: 4
Constraints:

1 <= nums.length <= 1000
0 <= nums[i] <= 10^6
1 <= m <= min(50, nums.length)
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	m := 2
	fmt.Printf("result:%d\n ", splitArray(nums, m))

}

func splitArray(nums []int, m int) int {
	left, right := 0, 0

	for _, num := range nums {
		right += num

		if left < num {
			left = num
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if isValid(nums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// Check if it is possible to split the array into m subarrays
// such that the sum of each subarray is less than or equal to target

func isValid(nums []int, m, target int) bool {
	count := 1
	sum := 0

	for _, num := range nums {
		sum += num

		if sum > target {
			count++
			sum = num

			if count > m {
				return false
			}
		}
	}
	return true

}
