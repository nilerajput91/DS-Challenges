/* The Valid Sudoku problem is a classic problem that involves checking whether a given 9x9 Sudoku board is valid. Here's the problem statement:

Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
*/

package main

import "fmt"

func main() {
	sudoku := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("sudko board:")

	for _, row := range sudoku {
		fmt.Println(string(row))
	}

	isValid := isValidSudko(sudoku)
	fmt.Println("Is valid:", isValid)

}

func isValidSudko(board [][]byte) bool {

	rowSets := make([]map[byte]bool, 9)
	columnSet := make([]map[byte]bool, 9)
	boxSet := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowSets[i] = make(map[byte]bool)
		columnSet[i] = make(map[byte]bool)
		boxSet[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				digit := board[i][j] - '0' // Convert byte to int

				if rowSets[i][digit] || columnSet[j][digit] || boxSet[(i/3)*3+j/3][digit] {
					return false
				}

				//add the digits sets

				rowSets[i][digit] = true
				columnSet[j][digit] = true
				boxSet[(i/3)*3+j/3][digit] = true
			}

		}
	}
	return true
}
