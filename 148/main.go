/* The Trapping Rain Water problem is a classic algorithmic problem where you are given an array representing the heights of walls, and you need to calculate how much rainwater can be trapped between the walls.

Here's the basic idea to solve this problem:

Iterate through the array to find the highest wall and its index.
Split the array into two parts at the highest wall's index.
For each part, iterate from left to right (or right to left) and keep track of the highest wall on the left (or right).
For each position, calculate the amount of water that can be trapped based on the height of the current wall and the highest wall on its left and right sides.
*/

package main

import "fmt"

func main() {
	input1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("triped water:%v\n", trap(input1))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	totalWater := 0
	maxIndex := 0

	// Find the index of the highest wall

	for i := 1; i < len(height); i++ {
		if height[i] > height[maxIndex] {
			maxIndex = i
		}

	}

	// Process the left part of the highest wall

	leftMax := 0

	for i := 0; i < maxIndex; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		} else {
			totalWater += leftMax - height[i]
		}
	}

	// Process the right part of the highest wall
	rightMax := 0

	for i := 0; i < maxIndex; i++ {
		if height[i] > rightMax {
			rightMax = height[i]
		} else {
			totalWater += rightMax - height[i]
		}
	}

	return totalWater

}
