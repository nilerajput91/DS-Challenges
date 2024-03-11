// Matrix in Snake Pattern

package main

import "fmt"

func main() {

	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rows, colum := len(mat), len(mat[0])

	for i := 0; i < rows; i++ {
		if i%2 == 0 {
			for j := 0; j < colum; j++ {
				fmt.Print(mat[i][j], " ")
			}
		} else {
			for j := colum - 1; j >= 0; j-- {
				fmt.Print(mat[i][j], " ")

			}
		}
		fmt.Println()

	}
}
