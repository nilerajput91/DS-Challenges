//Longest Even Odd Subarray

package main

import "fmt"

func main() {
	arr := []int{5, 10, 20, 6, 3, 8}
	n := 6

	fmt.Printf("maXEvenOdd:%v\n", maxEvenOdd(arr, n))

}

func maxEvenOdd(arr []int, n int) int {
	res := 1
	curr := 1

	for i := 1; i < n; i++ {
		if (arr[i]%2 == 0 && arr[i-1]%2 != 0) ||
			(arr[i]%2 != 0 && arr[i-1]%2 == 0) {

			curr++
			res = max(res, curr)
		} else {
			curr = 1
		}
	}
	return res
}
