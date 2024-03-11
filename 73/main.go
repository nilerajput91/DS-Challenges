package main

import "fmt"

func main() {

	num := 873

	result := getSum1(num)
	fmt.Println(result)

}

func getSum1(num int) int {
	result := 0
	for num > 0 {
		result = result + num%10
		num = num / 10
	}

	return result
}
