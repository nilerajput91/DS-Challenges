package main

import (
	"fmt"
	"sort"
)

func main() {
		// Example 1
		arr1 := []int{2, 4, 1, 3, 5}
		n1 := 5
		result1 := thirdlargestEle(arr1, n1)
		fmt.Println(result1) // Output: 3

}

func thirdlargestEle(arr []int, n int) int {

	if n < 3 {
		return -1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

	return arr[3]
}
