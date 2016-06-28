package main

import (
	"./graph"
	"fmt"
)

// Finds the cheapest path between two vertices in a graph
// Returns a slice of the steps and the total cost of those steps
func Dijkstra(g *graph.Graph, origin *graph.Vertex, destination *graph.Vertex) ([]*graph.Vertex, int) {
	// Add all the vertices to a set
	vertices := graph.NewVertexSet()
	for _, vertex := range g.Vertices {
		vertices.Add(vertex)
	}

	// We start at origin, so it will be selected first
	origin.Cost = 0

	for vertices.HasContent() {
		currentVertex := vertices.PopCheapest()

		if currentVertex == destination {
			break
		}

		// For each neighbour of currentVertex where the neighbour is still unvisited
		for _, edge := range currentVertex.Edges {
			neighbour := edge.Destination
			if !vertices.Contains(neighbour) {
				continue
			}

			cost := currentVertex.Cost + edge.Weight
			if cost < neighbour.Cost {
				neighbour.Cost = cost
				neighbour.Prev = currentVertex
			}
		}

	}

	return destination.Path(), destination.Cost
}

func main() {
	g := graph.GenerateGraph(10, 3.0, 20)

	g.Display()

	fmt.Println("Finding path between A and C...")
	origin := g.FindVertex("A")
	destination := g.FindVertex("C")
	path, cost := Dijkstra(g, origin, destination)

	fmt.Println("Result:")
	for _, vertex := range path {
		fmt.Printf(" -> %v", vertex.Id)
	}
	fmt.Printf(" (%v)\n", cost)
}
