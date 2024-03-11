package main

import "fmt"

func main() {
	arr := []int{10, 20, 10, 5, 15}

	fmt.Printf("CalPreFixSum is :%v\n", calpreFixSum(arr))
}

func calpreFixSum(arr []int) []int {
	n := len(arr)

	prefixSum := make([]int, n)

	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum

}
