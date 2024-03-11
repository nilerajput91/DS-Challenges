package main

import "fmt"

func main() {
	arr2 := []int{1, 2, 2, 4,2,2}
	n2 := 4
	result2 := removeDuplicate(arr2, n2)
	fmt.Println(result2)        // Output: 3
	fmt.Println(arr2[:result2]) // Output: [1 2 4]

}

func removeDuplicate(arr []int, n int) int {
	if n == 0 {
		return 0
	}

	count := 1

	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			fmt.Println(arr[i-1])
			arr[count] = arr[i]
			count++

		}
	}

	return count

}
