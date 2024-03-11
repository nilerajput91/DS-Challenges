//Leaders in an Array problem

package main

import "fmt"

func main() {
	arr := []int{7, 10, 4, 10, 6, 5, 2}

	leader(arr)
}

func leader(arr []int) {
	curr_leader := arr[len(arr)-1]
	fmt.Printf("Current Leader:%d\n", curr_leader)

	for i := len(arr) - 2; i >= 0; i-- {
		if curr_leader < arr[i] {
			curr_leader = arr[i]
			fmt.Printf("Current Leader:%d\n", curr_leader)

		}
	}

}
