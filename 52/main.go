package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	LrotateOne(arr)
	fmt.Println(arr)

}

// LrotateOne rotates the elements of the array arr to the left by one position.
func LrotateOne(arr []int) {
	// Store the first element of the array in a temporary variable.
	temp := arr[0]

	// Shift all elements to the left by one position.
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}

	// Place the stored first element at the end of the array.
	arr[len(arr)-1] = temp
}
