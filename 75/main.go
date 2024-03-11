package main

import "fmt"

func main() {

	input := []int{1, 5, 11, 5}

	isBool, result := partationEqulSubset(input)

	fmt.Printf("bool:%v,result:%d \n", isBool, result)

}

func partationEqulSubset(arr []int) (bool, int) {
	totalsum := 0

	for i := 0; i < len(arr); i++ {
		totalsum += i

	}

	if totalsum%2 != 0 {
		return false, 0
	}

	targetedSum := totalsum / 2

	memo := make(map[int]bool)

	result := subsetSum(arr, len(arr)-1, targetedSum, memo)
	return result, targetedSum

}

func subsetSum(nums []int, index, targetSum int, memo map[int]bool) bool {

	if targetSum == 0 {
		return true
	}

	if index < 0 || targetSum < 0 {
		return false
	}

	if result, ok := memo[targetSum*1000+index]; ok {
		return result
	}

	included := subsetSum(nums, index-1, targetSum-nums[index], memo)

	excluded := subsetSum(nums, index-1, targetSum, memo)

	memo[targetSum*1000+index] = included || excluded

	return memo[targetSum*1000+index]

}
