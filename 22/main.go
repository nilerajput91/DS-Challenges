package main

import "fmt"

func main() {
	// Example 2
	arr2 := []int{3, 1, 3, 3}
	fmt.Println(finMajorityOfElement(arr2)) // Output: 3

}

func finMajorityOfElement(arr []int) int {

	n := len(arr)

	mapEle := make(map[int]int)

	for i := 0; i < n; i++ {
		mapEle[arr[i]]++

	}

	for key, val := range mapEle {
		if val > 1 {
			return key
		}

	}
	return -1
}
