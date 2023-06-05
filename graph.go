package graphs

import (
	"fmt"
	"strings"
)

type Node string

type Graph struct {
	Edges map[Node][]Node
}

func NewGraph() GraphProvider {
	return &Graph{
		Edges: make(map[Node][]Node),
	}
}

type GraphProvider interface {
	// AddEdge adds an edge between the source and destination nodes.
	AddEdge(source, destination Node)

	// IsEulerian checks if the graph is Eulerian.
	IsEulerian() bool

	// PrintGraph prints the graph.
	PrintGraph()
}

func (g *Graph) AddEdge(source, destination Node) {
	g.Edges[source] = append(g.Edges[source], destination)
	g.Edges[destination] = append(g.Edges[destination], source)
}

func (g *Graph) IsEulerian() bool {
	oddCount := 0
	for _, degree := range g.Edges {
		if len(degree)%2 != 0 {
			if oddCount > 2 {
				//  It cannot be Eulerian.
				break
			}

			oddCount++
		}
	}

	return oddCount == 0 || oddCount == 2
}

func (g *Graph) PrintGraph() {
	for node, neighbors := range g.Edges {
		str := strings.Builder{}
		for _, neighbor := range neighbors {
			str.WriteString(fmt.Sprintf("%s,", neighbor))
		}

		fmt.Printf("[node:%s][neighbors:%s]\n", node, strings.TrimSuffix(str.String(), ","))
	}
}
