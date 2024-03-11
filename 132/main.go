// Matrix boundary traversal
package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 2, 3}, {5, 6, 7}, {1, 2, 3}}
    fmt.Println("Input:")
    for _, row := range matrix1 {
        fmt.Println(row)
    }
    fmt.Println("Output:")
    boundryTraversal(matrix1)


}

func boundryTraversal(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, rows-1, 0, cols-1

	for top <= bottom && left <= right {
		//traverse top row
		for i := left; i <= right; i++ {
			fmt.Println(matrix[top][i], " ")
		}
		top++
		//treverse right column
		for i := top; i < bottom; i++ {
			fmt.Println(matrix[i][right], " ")
		}
		right--

		//traverse bottom row
		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Println(matrix[bottom][i], " ")
			}
			bottom--
		}

		//travrese left column
		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Println(matrix[i][left], " ")
			}
			left++
		}

	}

}
