package main

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	arr1 := []int{1, 2, 3, 2}
	n1, x1, y1 := 4, 1, 2
	fmt.Println(minDist(arr1, n1, x1, y1)) // Output: 1

	// Example 2
	arr2 := []int{86, 39, 90, 67, 84, 66, 62}
	n2, x2, y2 := 7, 42, 12
	fmt.Println(minDist(arr2, n2, x2, y2)) // Output: -1

}

func minDist(array []int, n, x, y int) int {
	lastX, lastY := -1, -1
	minDistance := math.MaxInt

	for i := 0; i < n; i++ {
		if array[i] == x {
			lastX = i
			if lastY != -1 {
				minDistance = min(minDistance, lastX-lastY)
			}
		} else if array[i] == y {
			lastY = i
			if lastX != -1 {
				minDistance = min(minDistance, lastY-lastX)
			}
		}
	}

	if lastX == -1 || lastY == -1 {
		return -1

	}

	return minDistance
}
