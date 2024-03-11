// Trapping Rain Water
package main

import "fmt"

func main() {
	arrWater := []int{3, 0, 2, 0, 4}
	n := 5

	fmt.Printf("water pour:%v\n", getWater(arrWater, n))

}

func getWater(arr []int, n int) int {
	res := 0

	lmax := make([]int, n)
	rmax := make([]int, n)

	lmax[0] = arr[0]

	for i := 1; i < n; i++ {
		lmax[i] = max(arr[i], lmax[i-1])
	}

	rmax[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(arr[i], rmax[i+1])

	}

	for i := 1; i < n-1; i++ {
		res = res + (min(lmax[i], rmax[i]) - arr[i])

	}
	return res

}
