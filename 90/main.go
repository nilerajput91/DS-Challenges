//Left Rotate an Array by D places

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	d := 2
	n := 5

	fmt.Println(leftRoatateByD(arr, n, d))

}

func leftRoatateByD(arr []int, n, d int) []int {
	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
	reverse(arr, 0, n-1)
	return arr

}

func reverse(arr []int, low, high int) {
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
}
