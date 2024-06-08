/* Given an array of integers nums and a positive integer k,
you need to divide nums into sets of k consecutive numbers.
Return true if you can divide nums into such sets, and false otherwise.
Input: nums = [1, 2, 3, 3, 4, 4, 5, 6], k = 4
Output: true
Explanation: The array can be divided into the sets [1, 2, 3, 4], [3, 4, 5, 6].
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 3, 3, 4, 4, 5, 6}
	k := 4
	fmt.Println(isPossibleDivide(nums, k))

}

func isPossibleDivide(nums []int, k int) bool {
	// Create a map to store the frequency of each number in nums

	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	// Get the keys (numbers) from the map and sort them

	keys := make([]int, 0, len(countMap))

	for key := range countMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		// Get the count of the current number

		count := countMap[key]
		// If the count is greater than 0, try to form a set of size k

		if count > 0 {
			for i := 1; i < k; i++ {
				next := key + i

				nextCount := countMap[next]
				// If there are not enough consecutive numbers to form a set of size k, return false

				if nextCount < count {
					return false
				}
				// Reduce the count of the consecutive number

				countMap[next] -= count
			}
		}
	}
	return true
}
