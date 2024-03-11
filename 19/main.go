package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr1 := []int{3, 30, 34, 5, 9}
	result1 := largestNumber(arr1)
	fmt.Println(result1) // Output: 9534330

}

func largestNumber(arr []int) string {

	sort.Slice(arr, func(i, j int) bool {
		str1 := strconv.Itoa(arr[i])
		str2 := strconv.Itoa(arr[j])
		return str1+str2 > str2+str1
	})

	result := ""
	for _, value := range arr {
		result += strconv.Itoa(value)
	}

	result = removeleadingZero(result)

	if result == "" {
		return "0"
	}

	return result
}

func removeleadingZero(input string) string {
	i := 0
	for i < len(input)-1 && input[i] == 0 {
		i++
	}
	return input[i:]
}
