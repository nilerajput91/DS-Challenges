//multidimential array in golang

package main

import "fmt"

func main() {
	//create a 2D array with 3 rows and 2 columns
	rows, cols := 3, 2

	arr := make([][]int, rows)

	for i := range arr {
		arr[i] = make([]int, cols)

	}

	arr[0][0] = 1
	arr[0][1] = 2
	arr[1][0] = 3
	arr[1][1] = 4
	arr[2][0] = 5
	arr[2][1] = 6

	//print the 2D array
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
