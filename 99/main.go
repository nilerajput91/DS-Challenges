// Majority Element

package main

import "fmt"

func main() {
	arr := []int{6,8,4,8,8}
	n := 5

	fmt.Printf("majority of element index is: %d\n", majorityOfElement(arr, n))

}

func majorityOfElement(arr []int, n int) int {

	for i := 0; i < n; i++ {
		count := 1

		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}

		}
		if count > n/2 {
			return i
		}

	}
	return -1

}
