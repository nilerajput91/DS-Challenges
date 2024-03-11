package main

import "fmt"

func main() {

	maze := [][]int{

		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
	}

	fmt.Println("Maze:")

	for _, rows := range maze {
		fmt.Println(rows)
	}

	solution:=solveMaze(maze)

	fmt.Println("Solution:",solution)



	printSolution(solution)

}

func solveMaze(maze [][]int) [][]int {
	rows := len(maze)

	if rows == 0 {
		return nil
	}
	cols := len(maze[0])

	solution := make([][]int, rows)

	for i, _ := range solution {
		solution[i] = make([]int, cols)
	}

	// initialize the solution matrix with zeros

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			solution[i][j] = 0
		}
	}

	if solveMazeUtil(maze, 0, 0, solution) {
		return solution
	}

	return nil

}

func solveMazeUtil(maze [][]int, x, y int, solution [][]int) bool {
	rows := len(maze)
	cols := len(maze[0])

	//chehck if the current cell is in valid position
	if x >= 0 && x < rows && y >= 0 && y < cols && maze[x][y] == 1 {
		//mark the current cell as part of ths soltuion
		solution[x][y] = 1

		//if reach to destination mark true
		if x == rows-1 && y == cols-1 {
			return true
		}
		//move down

		if solveMazeUtil(maze, x+1, y, solution) {
			return true
		}

		//move right

		if solveMazeUtil(maze, x, y+1, solution) {
			return true
		}

		// if none of the above moment work then  backtrack
		solution[x][y] = 0
		return false
	}
	return false

}

func printSolution(solution [][]int) {
	if solution == nil {
		fmt.Println("no solution exists")
		return
	}

	rows := len(solution)
	cols := len(solution[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d  ", solution[i][j])
		}
		fmt.Println()
	}
}
