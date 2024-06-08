package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperture(temperatures))

}

func dailyTemperture(T []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(T))

	for i, temp := range T {
		for len(stack) > 0 && temp > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex

		}
		stack = append(stack, i)
	}
	return result
}
