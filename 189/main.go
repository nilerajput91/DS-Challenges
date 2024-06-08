/* The "Number of Sub-arrays With Odd Sum" problem asks us to find the number of subarrays within a given array whose sum is an odd number. We need to return this count modulo 10^9 + 7.

Here's a formal problem statement:

Given an array of integers `arr`, return the number of subarrays with an odd sum.

A subarray is a contiguous subsequence of the array.

Example:

Input: arr = [1,3,5]

Output: 4

Explanation: The subarrays with an odd sum are [1], [3], [5], and [1,3,5].

Constraints:
- 1 <= arr.length <= 10^5
- 1 <= arr[i] <= 100
- arr[i] is even or odd. */

package main

import "fmt"

func main() {
	nums := []int{1, 2}
	fmt.Println("missing positive number is", firstMissingPositivie(nums))

}

func firstMissingPositivie(nums []int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		if num > 0 {
			numMap[num] = true
		}
	}
	for i := 1; i < len(nums); i++ {
		if !numMap[i] {
			// If not present, return the missing integer
			return i
		}
	}
	// If all integers from 1 to the length of the input array are present, return the next positive integer

	return len(nums) + 1

}
