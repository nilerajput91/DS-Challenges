//Sliding Window Technique

package main

import "fmt"

func main() {
	arr := []int{100, 200, 300, 400}
	n := 4
	k := 2

	fmt.Printf("maxSum:%d\n", maxSum(arr, n, k))

}

func maxSum(arr []int, n, k int) int {
	var result int
	var curr int

	for i := 0; i < k; i++ {
		curr += arr[i]
		result = curr

	}
	for i := k; i < n; i++ {
		curr = curr + arr[i] - arr[i-k]
		result = max(result, curr)
	}
	return result

}
