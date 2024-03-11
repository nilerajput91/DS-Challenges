//DFS using graphs

package main

import "fmt"

func main() {
	graph := NewGraph()
	graph.addEdage(0, 1)
	graph.addEdage(0, 2)
	graph.addEdage(1, 2)
	graph.addEdage(2, 3)
	graph.addEdage(3, 3)

	startVertex := 0
	fmt.Printf("DFS from vertex %d: ", startVertex)
	graph.DFS(startVertex)
	fmt.Println()

}

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (e *Graph) addEdage(u, v int) {
	e.adjacencyList[u] = append(e.adjacencyList[u], v)
	e.adjacencyList[v] = append(e.adjacencyList[v], u)
}

func (e *Graph) DFSutil(vertex int, visited map[int]bool) {
	visited[vertex] = true

	fmt.Printf("%d ", vertex)

	for _, nebhiour := range e.adjacencyList[vertex] {
		if !visited[nebhiour] {
			e.DFSutil(nebhiour, visited)
		}
	}

}

func (e *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)
	e.DFSutil(startVertex, visited)
}
