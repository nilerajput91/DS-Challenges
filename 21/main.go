package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{5, 2, 3, 6, 4, 4, 6, 6}
	fmt.Println(contigousElement(arr1)) // Output: Yes

}

func contigousElement(arr []int) string {
	n := len(arr)

	sort.Ints(arr)
	fmt.Println(arr)

	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1]+1 {
			return "NO"
		}
	}

	return "Yes"

}
