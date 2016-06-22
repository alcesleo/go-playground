package main

import (
	"fmt"
)

// Data structure of a weighted directed graph

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Id    string
	Edges []*Edge
}

type Edge struct {
	Destination *Vertex
	Weight      int
}

// Constructors

func NewWeightedGraph() *Graph {
	return &Graph{
		Vertices: make([]*Vertex, 0),
	}
}

func NewVertex(Id string) *Vertex {
	return &Vertex{
		Id:    Id,
		Edges: make([]*Edge, 0),
	}
}

func NewEdge(destination *Vertex, weight int) *Edge {
	return &Edge{
		Destination: destination,
		Weight:      weight,
	}
}

// Instance methods

func (g *Graph) AddVertex(vertex *Vertex) {
	g.Vertices = append(g.Vertices, vertex)
}

func (g *Graph) FindVertex(id string) *Vertex {
	for _, vertex := range g.Vertices {
		if vertex.Id == id {
			return vertex
		}
	}

	panic("No vertex found")
}

func (v *Vertex) AddEdge(edge *Edge) {
	v.Edges = append(v.Edges, edge)
}

func (v *Vertex) Connect(destination *Vertex, weight int) {
	edge := NewEdge(destination, weight)
	v.AddEdge(edge)
}

// Generate data

func GenerateGraph() *Graph {
	graph := NewWeightedGraph()

	// Create vertices
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")

	// Connect them to each other
	a.Connect(b, 5)
	a.Connect(d, 10)
	b.Connect(d, 2)
	b.Connect(c, 20)
	d.Connect(c, 10)

	// Add the vertices to the graph
	graph.AddVertex(a)
	graph.AddVertex(b)
	graph.AddVertex(c)
	graph.AddVertex(d)

	return graph
}

// Print graph

func (g *Graph) Display() {
	for _, vertex := range g.Vertices {
		fmt.Println(vertex.Id)
		for _, edge := range vertex.Edges {
			fmt.Printf("  -> %v (%v)\n", edge.Destination.Id, edge.Weight)
		}
		fmt.Println()
	}
}

// Find the path between them

// Finds the cheapest path between two vertices in a graph
// Returns a slice of the steps and the total cost of those steps
// When given debug=true also outputs logging of every path evaluated
func (graph *Graph) Dijkstra(origin *Vertex, destination *Vertex, debug bool) ([]*Vertex, int) {

}

func main() {
	graph := GenerateGraph()

	fmt.Println("Randomly generated graph:")
	graph.Display()
	fmt.Println()

	fmt.Println("Finding path between A and C:")

	origin := graph.FindVertex("A")
	destination := graph.FindVertex("C")
	path, cost := graph.Dijkstra(origin, destination, true)
}
