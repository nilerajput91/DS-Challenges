//Adjacency List implementation graph based questions

package main

import "fmt"

func main() {
	graph := NewGraph()
	graph.addEdage(0, 1)
	graph.addEdage(0, 2)
	graph.addEdage(1, 2)
	graph.addEdage(2, 3)
	graph.addEdage(3, 3)
	graph.printAdjuencyList()

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
	g.adjuencyList[v] = append(g.adjuencyList[v], u) //for undirected graph

}

func (g *Graph) printAdjuencyList() {
	for vertex, neighbours := range g.adjuencyList {
		fmt.Printf("vertex:%d\n", vertex)

		for _, neighbor := range neighbours {
			fmt.Printf("neighbours:%d\n", neighbor)
		}
		fmt.Println()
	}
}
