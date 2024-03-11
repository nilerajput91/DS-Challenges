package main

import "fmt"

func main() {
	arr := []int{10, 8, 0, 0, 12, 0}
	fmt.Printf("after moving zero:%v\n", moveZero(arr,6))

}

func moveZero(arr []int, n int) []int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[count], arr[i] = arr[i], arr[count]
			count++
		}
	}
	return arr

}
