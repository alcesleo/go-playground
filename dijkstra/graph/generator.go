package graph

import (
	"math/rand"
)

// Converts an int to a letter-id
func IntToId(n int) string {
	numLetters := 26
	asciiOffset := 64
	rest := n
	result := make([]byte, 0)

	for rest > 0 {
		temp := rest % numLetters
		if temp == 0 {
			temp = numLetters
		}
		rest = rest / numLetters
		char := byte(temp + asciiOffset)

		result = append([]byte{char}, result...)
	}
	return string(result[:])
}

func RandomVertex(g *Graph) *Vertex {
	numVertices := len(g.Vertices)
	if numVertices == 0 {
		panic("No vertices in graph")
	}
	return g.Vertices[rand.Intn(len(g.Vertices))]
}

// Generate data

// Generate a random graph with a given number of vertices, and a sparseness
// factor deciding how many edges should be created. A graph with 10 vertices
// and a sparseness factor of 2.0 will have around 20 edges.
func GenerateGraph(vertices int, sparseness float32, maxCost int) *Graph {
	graph := NewGraph()

	// Create the starting point so there's something to connect from in the loop
	graph.AddVertex(NewVertex(IntToId(1)))

	// Create a vertex with a connection from a random vertex already created
	for n := 2; n < vertices+1; n++ {
		vertex := NewVertex(IntToId(n))
		randomVertex := RandomVertex(graph)
		randomVertex.Connect(vertex, rand.Intn(maxCost))
		graph.AddVertex(vertex)
	}

	// Create an edge between 2 random vertices until the sparseness factor is reached
	edgesLeft := int((float32(vertices) * sparseness)) - vertices
	for n := 0; n < edgesLeft; n++ {
		source := RandomVertex(graph)

		// Make sure the destination is not the source
		destination := source
		for source == destination {
			destination = RandomVertex(graph)
		}

		source.Connect(destination, rand.Intn(maxCost))
	}

	return graph
}
