// Maximum subarray sum
package main

import "fmt"

func main() {
	arr := []int{1, -2, 3, -1, 2}
	fmt.Printf("maximum sub array:%v\n", maxSumSubarray(arr, 5))

}

func maxSumSubarray(arr []int, n int) int {
	res := arr[0]
	maxEnding := arr[0]

	for i := 1; i < n; i++ {
		maxEnding = max(maxEnding+arr[i], arr[i])
		res = max(res, maxEnding)

	}
	return res
}
