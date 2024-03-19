// Remove Duplicate from Sorted Array
package main

import "fmt"

func main() {

	nums := []int{2, 2, 3, 5}
	newLength := removeDuplicates(nums)
	fmt.Println("new length:", nums[:newLength])

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	nonDuplicateArray := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nonDuplicateArray] = nums[i]
			nonDuplicateArray++
		}
	}

	return nonDuplicateArray
}
