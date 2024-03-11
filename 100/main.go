//Minimum Consecutive Flips

package main

import "fmt"

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 0, 1, 0, 1}
	n := 10

	 printGroups(arr, n)

}

func printGroups(arr []int, n int) {
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			if arr[i] != arr[0] {
				fmt.Printf("From:%d to ", i)
			} else {
				fmt.Println(i - 1)
			}
		}
	}

	if arr[n-1] != arr[0] {
		println(n - 1)
	}
}
