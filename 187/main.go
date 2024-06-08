package main

import "fmt"

func main() {

	arr:=[]int{1,5,6,9,11,16}
	target:=11

	fmt.Printf("find %d",binarysearch(arr,target))

}

func binarysearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left<=right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1

}
