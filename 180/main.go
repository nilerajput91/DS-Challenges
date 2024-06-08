/* Here's the problem statement for the "Capacity To Ship Packages Within D Days" problem:

You have n packages that you need to ship to a destination. The ith package has a weight weights[i].
 You want to determine the minimum capacity of a ship that will allow you to ship all the packages within D days

The capacity of the ship is the maximum weight it can carry.
 The ship will be loaded with packages in the order they are given. Each day, 
 the ship will be loaded with as many packages as possible while respecting the weight limit.

Return the minimum capacity of the ship that will allow you to ship all the packages within D days.

Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
Output: 15
Explanation: A ship capacity of 15 allows you to ship all packages in 5 days:
- Day 1: 1, 2, 3, 4, 5
- Day 2: 6, 7
- Day 3: 8
- Day 4: 9
- Day 5: 10

Note that the ship capacity provided as an answer is the minimum such capacity that allows you to ship all packages in D days, and it can be proven that it is the only solution possible.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5

	fmt.Println(shipWithDays(weights, D))

}

func shipWithDays(weights []int, D int) int {
	low, high := 0, math.MaxInt32

	for low < high {
		mid := low + (high-low)/2

		if canShip(weights, D, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canShip(weights []int, D, capacity int) bool {
	days := 1
	curWeight := 0

	for _, weight := range weights {
		if curWeight+weight > capacity {
			days++
			curWeight = 0

		}

		curWeight += weight
	}
	return days <= D
}
