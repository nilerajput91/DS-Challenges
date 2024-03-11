//Maximum consecutive 1s

package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 0, 1, 1}
	n := 6
	fmt.Printf("max consecutive element is :%v\n", maxConsecutive1s(arr, n))

}

func maxConsecutive1s(arr []int, n int) int {
	res := 0
	var curr int
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			curr = 0
		} else {
			curr++
			res = max(res, curr)

		}
	}
	return res
}
