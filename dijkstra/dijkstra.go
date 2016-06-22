package main

import (
	"fmt"
	"math"
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

// Helper type for the algorithm, Go unfortunately doesn't have Set yet :(
type VertexSet struct {
	Vertices map[*Vertex]bool
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

func NewVertexSet() *VertexSet {
	return &VertexSet{
		make(map[*Vertex]bool),
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

func (vs *VertexSet) Add(vertex *Vertex) {
	vs.Vertices[vertex] = true
}

func (vs *VertexSet) Has(vertex *Vertex) bool {
	_, exists := vs.Vertices[vertex]
	return exists
}

func (vs *VertexSet) Remove(vertex *Vertex) {
	delete(vs.Vertices, vertex)
}

func (vs *VertexSet) Len() int {
	return len(vs.Vertices)
}

// Generate data

func GenerateGraph() *Graph {
	graph := NewWeightedGraph()

	// Create vertices
	a := NewVertex("A")
	b := NewVertex("B")
	c := NewVertex("C")
	d := NewVertex("D")
	e := NewVertex("E")
	f := NewVertex("F")
	g := NewVertex("G")

	// Connect them to each other
	a.Connect(b, 5)
	a.Connect(d, 10)
	a.Connect(f, 50)
	b.Connect(d, 2)
	b.Connect(c, 20)
	d.Connect(c, 10)
	d.Connect(e, 1)
	e.Connect(g, 1)
	e.Connect(c, 5)
	g.Connect(f, 2)
	f.Connect(c, 1)

	// Add the vertices to the graph
	graph.AddVertex(a)
	graph.AddVertex(b)
	graph.AddVertex(c)
	graph.AddVertex(d)
	graph.AddVertex(e)
	graph.AddVertex(f)
	graph.AddVertex(g)

	return graph
}

// Print graph

func (g *Graph) Display() {
	for _, vertex := range g.Vertices {
		fmt.Printf(" - %v\n", vertex.Id)
		for _, edge := range vertex.Edges {
			fmt.Printf("   |-> %v (%v)\n", edge.Destination.Id, edge.Weight)
		}
		fmt.Println()
	}
}

// Find the path between them

// Finds the cheapest path between two vertices in a graph
// Returns a slice of the steps and the total cost of those steps
// When given debug=true also outputs logging of every path evaluated
func (graph *Graph) Dijkstra(origin *Vertex, destination *Vertex) ([]*Vertex, int) {
	unvisited := NewVertexSet()
	cost := make(map[*Vertex]int)     // Maps a vertex to its current calculated cost
	prev := make(map[*Vertex]*Vertex) // Maps a vertex to the previous vertex in the cheapest path to get here

	// Set up all the maps keeping track of values
	for _, vertex := range graph.Vertices {
		cost[vertex] = math.MaxInt32
		prev[vertex] = nil
		unvisited.Add(vertex)
	}

	// We start at origin, so it will be selected first
	cost[origin] = 0

	// While unvisited is not empty
	for unvisited.Len() > 0 {
		// Find the vertex with the lowest cost
		// TODO: Extract this to VertexPriorityQueue
		var currentVertex *Vertex = nil
		for vertex := range unvisited.Vertices {
			if cost[currentVertex] > cost[vertex] || currentVertex == nil {
				currentVertex = vertex
			}
		}

		// If we are at the destination we are done
		if currentVertex == destination {
			break
		}

		// We have now visited this vertex
		unvisited.Remove(currentVertex)

		// For each neighbour of currentVertex where the neighbour is still unvisited
		for _, edge := range currentVertex.Edges {
			neighbour := edge.Destination
			if !unvisited.Has(neighbour) {
				continue
			}

			alt := cost[currentVertex] + edge.Weight
			if alt < cost[neighbour] {
				cost[neighbour] = alt
				prev[neighbour] = currentVertex
			}
		}

	}

	// Extract the best path
	result := make([]*Vertex, 0)
	target := destination
	result = append(result, target) // Start from the target
	for prev[target] != nil {
		result = append([]*Vertex{prev[target]}, result...)
		target = prev[target]
	}
	return result, cost[destination]
}

func main() {
	graph := GenerateGraph()

	fmt.Println("Randomly generated graph:")
	graph.Display()
	fmt.Println()

	fmt.Println("Finding path between A and C...")
	origin := graph.FindVertex("A")
	destination := graph.FindVertex("C")
	path, cost := graph.Dijkstra(origin, destination)

	fmt.Println("Result:")
	for _, vertex := range path {
		fmt.Printf(" -> %v", vertex.Id)
	}
	fmt.Printf(" (%v)", cost)
}
