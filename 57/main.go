package main

import "fmt"

func main() {

	str := "aabc"

	result := leftMost(str)
	fmt.Println(result)

}

func leftMost(str string) int {

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return i
			}
		}
	}

	return -1
}
