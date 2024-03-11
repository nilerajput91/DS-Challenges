// Detect Cycle in Undirected Graph
package main

import "fmt"

func main() {
	graph := NewGraph(5)
	graph.AddEdge(1, 0)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(3, 4)

	if graph.IsCyclic() {
		fmt.Println("Graph contains cycles")
	}else{
		fmt.Println("Graph doesn't contain cycle")
	}

}

type Graph struct {
	Vertex int
	Adj    [][]int
}

// NewGraph creates a new instance of Graph with V vertices
func NewGraph(V int) *Graph {
	return &Graph{
		Vertex: V,
		Adj:    make([][]int, V),
	}
}

// AddEdge adds an undirected edge between vertices u and v

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u)

}

// isCyclicUtil is a recursive utility function to detect cycle in the graph
func (g *Graph) isCyclicUtil(v, parent int, visited []bool) bool {
	visited[v] = true

	for _, i := range g.Adj[v] {
		// If the adjacent vertex is not visited, then recur for it
		if !visited[i] {
			if g.isCyclicUtil(i, v, visited) {
				return true
			}
		} else if i != parent {
			// If an adjacent vertex is visited and not the parent of the current vertex,
			// then there is a cycle in the graph
			return true
		}

	}
	return false

}

func (g *Graph) IsCyclic() bool {
	visited := make([]bool, g.Vertex)

	for i := 0; i < g.Vertex; i++ {
		if !visited[i] {
			if g.isCyclicUtil(i, -1, visited) {

				return true
			}

		}
	}
	return false
}
