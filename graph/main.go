package main

import (
	"fmt"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g Graph) Print() {
	for k, v := range g.adjacencyList {
		fmt.Printf("%d -> %v\n", k, v)
	}
}

type GraphAdjMatrix struct {
	adjacencyMatrix [][]bool
	numNodes        int
}

func NewGraphAdjMatrix(numNodes int) *GraphAdjMatrix {

	matrix := make([][]bool, numNodes)

	for i, _ := range matrix {
		matrix[i] = make([]bool, numNodes)
	}

	return &GraphAdjMatrix{
		adjacencyMatrix: matrix,
		numNodes:        numNodes,
	}
}

func (g GraphAdjMatrix) Print() {
	for i := 0; i < g.numNodes; i++ {
		for j := 0; j < g.numNodes; j++ {
			fmt.Print(g.adjacencyMatrix[i][j], "\n")
		}
	}
}

func main() {
	g := NumberOfProvince([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})

	fmt.Println(g)
}
