package main

import "fmt"

func main() {
	nums:=[]int{1,3,5}
	fmt.Printf("subarrays:%d\n",numOfSubarrays(nums))

}

func numOfSubarrays(arr []int) int {
	// Initialize variables to keep track of the count of odd and even sums,
	// as well as the result count.
	oddSumCounts, evenSumCounts, result := 0, 0, 0

	sum := 0

	for _, num := range arr {
		sum += num
		// If the sum is odd, increment the odd sum count and add the even sum count to the result.
		if sum%2 == 1 {
			result += evenSumCounts + 1
			oddSumCounts++
			// If the sum is even, increment the even sum count and add the odd sum count to the result.
		} else {
			result += oddSumCounts
			evenSumCounts++
		}
	}
	// Return the result count modulo 10^9 + 7 as the answer can be large.
	return result % (1e9 + 7)

}
