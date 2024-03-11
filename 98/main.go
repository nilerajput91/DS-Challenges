//Maximum Circular Sum Subarray

package main

import "fmt"

func main() {
	arr := []int{5, -2, 3, 4}
	n := 4
	fmt.Printf("Max Subarray in circular is :%v\n", maxSubarrSumCircular(arr, n))

}

func maxSubarrSumCircular(arr []int, n int) int {
	res := arr[0]
	for i := 0; i < n; i++ {
		curr_max := arr[i]
		curr_sum := arr[i]

		for j := 1; j < n; j++ {
			index := (i + j) % n
			curr_sum += arr[index]
			curr_max = max(curr_max, curr_max)
		}

		res = max(res, curr_max)
	}

	return res

}
