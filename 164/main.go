/* Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

A region is surrounded by 'X' if all the cells in the region are either bordering
the board or are connected to an 'O' on the border by a path of 'O's.


X X X X
X O O X
X X O X
X O X X

After running the function, the board should be:

X X X X
X X X X
X X X X
X O X X

*/

package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)

}

func solve(board [][]byte) {
	if len(board) == 0 {

		return
	}
	rows, cols := len(board), len(board[0])

	// Mark 'O's connected to the border as safe by doing a DFS
	for i := 0; i < rows; i++ {
		if board[i][0] == '0' {
			//start dfs from the left border
			dfs(board, i, 0)
		}

		if board[i][cols-1] == '0' {
			//start dfs from the right border
			dfs(board, i, cols-1)
		}
	}

	for j := 0; j < cols; j++ {
		if board[0][j] == '0' {
			// Start DFS from the top border
			dfs(board, 0, j)

		}
		if board[rows-1][j] == '0' {
			// Start DFS from the bottom border
			dfs(board, rows-1, j)
		}
	}

	// Convert remaining 'O's to 'X's and revert safe 'O's back to 'O's
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '0' {
				board[i][j] = 'X'
			} else if board[i][j] == 'S' {
				board[i][j] = '0'

			}

		}
	}

}

func dfs(board [][]byte, i, j int) {

	if i <= 0 || j <= 0 || i >= len(board) || j >= len(board) || board[i][j] != '0' {
		return
	}
	// Mark the current 'O' as safe
	board[i][j] = 'S'

	// Perform DFS in all four directions
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)

	fmt.Println(board)

}
