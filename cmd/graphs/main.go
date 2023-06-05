package main

import (
	"fmt"

	"graphs"
)

func main() {
	graph := graphs.NewGraph()

	// Seven Bridges of Königsberg
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "A")
	graph.AddEdge("A", "C")
	graph.AddEdge("C", "A")
	graph.AddEdge("A", "D")
	graph.AddEdge("D", "C")
	graph.AddEdge("D", "B")

	graph.PrintGraph()

	fmt.Printf("[is_eulerian:%t]", graph.IsEulerian())
}
