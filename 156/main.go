/* Next Greater Element II is a problem where
you're given a circular array (the next element of the last element is the first element of the array) and
you need to find the Next Greater Number for every element. The Next Greater Number of a number x is the
first greater number to its traversing-order next in the array, which means you could search circularly to
find its next greater number.
 If it doesn't exist, output -1 for this number. */

package main

import "fmt"

func main() {

	Input := []int{1, 2, 1}
	fmt.Println(nextGreaterElement(Input))

}

func nextGreaterElement(nums []int) []int {
	n := len(nums)

	result := make([]int, n)
	stack := make([]int, 0)

	// Traverse the array twice to handle the circular nature of the array
	for i := 2*n - 1; i >= 0; i-- {
		/* Process the stack until it is empty or the current element
		is not greater than the top element
		 in the stack */
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i%n] {
			stack = stack[:len(stack)-1]  // Pop the top element from the stack

		}
		// If the stack becomes empty, there is no greater element to the right of the current element
		if len(stack) == 0 {
			result[i%n] = -1
		} else {
			// The next greater element is the element at the top of the stack
			result[i%n] = nums[stack[len(stack)-1]]
		}
		// Push the current element's index to the stack
		stack = append(stack, i%n)

	}
	return result

}
