package main

import "fmt"

func main() {
	fmt.Println(romanToIntNum("III")) // Output: 3

}

var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToIntNum(s string) int {
	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanValues[s[i]]
		// If the current value is less than the previous value,
		// subtract it from the total.
		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}
	return total
}
