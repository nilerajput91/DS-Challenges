// Maximum Difference Problem with Order
package main

import "fmt"

func main() {
	arr := []int{2, 3, 10, 6, 4, 8, 1}

	fmt.Printf("max difference:%v\n", maxDiff(arr))

}

func maxDiff(arr []int) int {
	result := arr[1] - arr[0]
	minVal := arr[0]

	for j := 1; j < len(arr); j++ {
		result = max(result, arr[j]-minVal)
		minVal = min(minVal, arr[j])
	}
	return result
}
