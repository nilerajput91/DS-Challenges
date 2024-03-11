package main

import "fmt"

func main() {
	arr := []int{2, 8, 30, 15, 25, 12}
	ceilingOnLeft := findingCellingLeft(arr)

	fmt.Println("Array:", arr)
	fmt.Println("Ceiling on Left for Each Element:", ceilingOnLeft)

}

func findingCellingLeft(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {

		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] {
			fmt.Println(arr[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
			fmt.Println(stack)


		}

		if len(stack) > 0 {
			result[i] = arr[stack[len(stack)-1]]
		} else {
			result[i] = -1
		}

		stack = append(stack, i)
	}
	return result
}
