// Prim's algorithm is used to find the minimum spanning tree (MST) of a connected, undirected graph.
package main

import (
	"fmt"
	"math"
)

func main() {

	graph := [][]int{{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0}}

	primMST(graph)

}

const V = 5

// findMinKey finds the vertex with the minimum key value, from the set of vertices not yet included in MST
func findMinKey(key []int, mstSet []bool) int {
	min := math.MaxInt32
	minIndex := -1
	for v := 0; v < V; v++ {
		if mstSet[v] == false && key[v] < min {
			min = key[v]
			minIndex = v
		}

	}
	return minIndex
}

func printMST(parent []int, graph [][]int) {
	fmt.Println("Edge \t Weight")
	for i := 1; i < V; i++ {
		fmt.Println(parent[i], "-", i, "\t", graph[i][parent[i]])
	}

}

// primMST finds the MST for a graph represented using adjacency matrix representation
func primMST(graph [][]int) {
	parent := make([]int, V)  // Array to store constructed MST
	key := make([]int, V)     // Key values used to pick minimum weight edge in cut
	mstSet := make([]bool, V) // To represent set of vertices included in MST

	// Initialize all keys as INFINITE
	for i := 0; i < V; i++ {
		key[i] = math.MaxInt32
		mstSet[i] = false
	}
	// Always include first 1st vertex in MST.
	key[0] = 0     // Make key 0 so that this vertex is picked as first vertex
	parent[0] = -1 // First node is always root of MST

	// The MST will have V vertices
	for count := 0; count < V-1; count++ {
		// Pick the minimum key vertex from the set of vertices not yet included in MST
		u := findMinKey(key, mstSet)
		// Add the picked vertex to the MST Set
		mstSet[u] = true

		// Update key value and parent index of the adjacent vertices of the picked vertex.
		// Consider only those vertices which are not yet included in MST
		for v := 0; v < V; v++ {
			// graph[u][v] is non-zero only for adjacent vertices of m
			// mstSet[v] is false for vertices not yet included in MST
			// Update the key only if graph[u][v] is smaller than key[v]
			if graph[u][v] != 0 && mstSet[v] == false && graph[u][v] < key[v] {
				parent[v] = u
				key[v] = graph[u][v]
			}

		}
	}
	printMST(parent, graph)

}
