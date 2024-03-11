/*
You are given two four digit prime numbers num1 and num2. Find the distance of the shortest path from Num1 to Num2 that can be attained by altering only single digit at a time such that every number that we get after changing a digit is a four digit prime number with no leading zeros.

Example 1:

Input:
num1 = 1033
num2 = 8179
Output: 6
Explanation:
1033 -> 1733 -> 3733 -> 3739 -> 3779 -> 8779 -> 8179.
There are only 6 steps reuired to reach num2 from num1.
and all the intermediate numbers are 4 digit prime numbers.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 1033
	num2 := 8179
	result := shortestPath(num1, num2)
	fmt.Printf("Example 1: Shortest path from %d to %d is %d\n", num1, num2, result)

}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false

		}

	}
	return true

}

func shortestPath(num1, num2 int) int {
	//create queue for bfs queue
	queue := []int{num1}

	visited := make(map[int]bool)
	visited[num1] = true

	distance := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == num2 {
				return distance
			}

			// explore all possible num recheable by changing a singel digit

			digit := 1

			for place := 1000; place > 0; place /= 10 {
				for replacement := 0; replacement <= 9; replacement++ {
					next := current - digit*(current/place%10) + digit*replacement

					//check if the next num is a four digit prime with no leading zero
					if !visited[next] && isPrime(next) {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				digit *= 10
			}

		}
		distance++

	}
	return -1
}
