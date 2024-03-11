// Detect Cycle in directed Graph

package main

import "fmt"

func main() {

	graph := NewGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	if graph.IsCyclic() {
		fmt.Println("graph contanis cyclic")
	} else {
		fmt.Println("not contains cyclic")
	}

}

type Graph struct {
	Vertex int
	Adj    [][]int
}

func NewGraph(V int) *Graph {
	return &Graph{
		Vertex: V,
		Adj:    make([][]int, V),
	}
}

// AddEdge adds a directed edge from u to v
func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
}

// isCyclicUtil is a recursive utility function to detect cycle in the graph
func (g *Graph) isCyclicUtil(v int, visited, recStack []bool) bool {
	if !visited[v] {
		visited[v] = true
		recStack[v] = true
	}
	// Recur for all the vertices adjacent to this vertex
	for _, i := range g.Adj[v] {
		if !visited[i] && g.isCyclicUtil(i, visited, recStack) {
			return true
		} else if recStack[i] {
			return true
		}

	}
	recStack[v] = false
	return false
}

// IsCyclic checks if the graph contains any cycle
func (g *Graph) IsCyclic() bool {
	visited := make([]bool, g.Vertex)
	recStack := make([]bool, g.Vertex)

	for i := 0; i < g.Vertex; i++ {
		if g.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}
	return false

}
