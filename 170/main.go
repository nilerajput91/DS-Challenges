/* Koko loves to eat bananas. There are N piles of bananas, the i-th pile has piles[i] bananas. \
The guards have gone and will come back in H hours.

Koko can decide her bananas-per-hour eating speed of K. Each hour, s
he chooses some pile of bananas and eats K bananas from that pile.
If the pile has less than K bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards return.

Return the minimum integer K such that she can eat all the bananas within H hours.

Example 1:
Input: piles = [3,6,7,11], H = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], H = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], H = 6
Output: 23
Constraints:

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9
Approach: */

package main

import "fmt"

func main() {

	piles := []int{30, 11, 23, 4, 20}
	H := 5

	fmt.Printf("eatingSpeed:%d\n", minEatingSpeed(piles, H))

}

// minEatingSpeed calculates the minimum eating speed K such that Koko can eat all the bananas within H hours.

func minEatingSpeed(piles []int, H int) int {
	// Initialize the search range from 1 to the maximum number of bananas in any pilel
	left, right := 1, getMax(piles)

	// Binary search to find the minimum eating speed

	for left < right {
		mid := left + (right-left)/2
		// If it is possible to eat all the bananas within H hours at speed mid, reduce the upper bound of the search range
		if canEatAll(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left

}

// getMax finds the maximum number of bananas in any pile

func getMax(piles []int) int {
	max := 0

	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	return max

}

// canEatAll checks if it is possible to eat all the bananas within H hours at speed K
func canEatAll(piles []int, k, h int) bool {
	hours := 0

	for _, pile := range piles {
		// Calculate the total hours required to eat all the bananas at speed K

		hours += (pile + k - 1) / k
	}
	// Return true if the total hours is less than or equal to H

	return hours <= h
}
