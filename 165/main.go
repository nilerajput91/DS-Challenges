/* Two Sum Problem: Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target */

package main

import "fmt"

func main() {
	nums := []int{2, 7, 8, 9}
	target := 9

	fmt.Printf("result:%v", twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {

		complement := target - num

		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i

	}
	return nil

}
