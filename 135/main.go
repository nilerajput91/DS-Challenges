// BFS search in graph
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
	fmt.Printf("BFS from vertex %d: ", startVertex)
	graph.BFS(startVertex)
	fmt.Println()

}

type Graph struct {
	adjuencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjuencyList: make(map[int][]int),
	}
}

func (g *Graph) addEdage(u, v int) {
	g.adjuencyList[u] = append(g.adjuencyList[u], v)
	g.adjuencyList[v] = append(g.adjuencyList[v], u)
}

func (g *Graph) BFS(startVertex int) {
	visited := make(map[int]bool)
	queue := []int{startVertex}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			fmt.Printf("%d ", vertex)
			visited[vertex] = true

			for _, neighbor := range g.adjuencyList[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}

		}
	}

}
