package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Data structure of a weighted directed graph

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Id    string
	Edges []*Edge

	// Used by the Dijkstra algorithm, could be extracted
	Cost int
	Prev *Vertex
}

type Edge struct {
	Destination *Vertex
	Weight      int
}

type VertexSet struct {
	Vertices map[*Vertex]bool
}

// Constructors

func NewGraph() *Graph {
	return &Graph{
		Vertices: make([]*Vertex, 0),
	}
}

func NewVertex(Id string) *Vertex {
	return &Vertex{
		Id:    Id,
		Edges: make([]*Edge, 0),

		Cost: math.MaxInt32,
		Prev: nil,
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

func (g *Graph) RandomVertex() *Vertex {
	numVertices := len(g.Vertices)
	if numVertices == 0 {
		panic("No vertices in graph")
	}
	return g.Vertices[rand.Intn(len(g.Vertices))]
}

func (g *Graph) Display() {
	for _, vertex := range g.Vertices {
		fmt.Printf(" - %v\n", vertex.Id)
		for _, edge := range vertex.Edges {
			fmt.Printf("   |-- %v --> %v\n", edge.Weight, edge.Destination.Id)
		}
		fmt.Println()
	}
}

func (v *Vertex) AddEdge(edge *Edge) {
	v.Edges = append(v.Edges, edge)
}

func (v *Vertex) Connect(destination *Vertex, weight int) {
	edge := NewEdge(destination, weight)
	v.AddEdge(edge)
}

// Follows Prev to the origin and returns a slice of all the steps from origin to here
func (v *Vertex) Path() []*Vertex {
	target := v
	result := []*Vertex{v} // Start from here

	for target.Prev != nil {
		result = append([]*Vertex{target.Prev}, result...)
		target = target.Prev
	}
	return result
}

func (v *Vertex) DisplayPath() {
	for _, vertex := range v.Path() {
		fmt.Printf(" -> %v", vertex.Id)
	}
	fmt.Printf(" (%v)\n", v.Cost)
}

func (vs *VertexSet) Add(vertex *Vertex) {
	vs.Vertices[vertex] = true
}

func (vs *VertexSet) Contains(vertex *Vertex) bool {
	_, exists := vs.Vertices[vertex]
	return exists
}

func (vs *VertexSet) Remove(vertex *Vertex) {
	delete(vs.Vertices, vertex)
}

func (vs *VertexSet) Len() int {
	return len(vs.Vertices)
}

func (vs *VertexSet) PopCheapest() *Vertex {
	var result *Vertex = nil
	for vertex := range vs.Vertices {
		if result == nil || result.Cost > vertex.Cost {
			result = vertex
		}
	}
	vs.Remove(result)

	return result
}

func (vs *VertexSet) HasContent() bool {
	return vs.Len() > 0
}

// Converts an int to a letter-id
func IntToId(n int) string {
	numLetters := 26
	asciiOffset := 64
	rest := n
	result := make([]byte, 0)

	for rest > 0 {
		letterPos := rest % numLetters
		if letterPos == 0 {
			letterPos = numLetters
		}
		rest = rest / numLetters
		char := byte(letterPos + asciiOffset)

		result = append([]byte{char}, result...)
	}
	return string(result[:])
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
		randomVertex := graph.RandomVertex()
		randomCost := rand.Intn(maxCost) + 1
		randomVertex.Connect(vertex, randomCost)
		graph.AddVertex(vertex)
	}

	// Create an edge between 2 random vertices until the sparseness factor is reached
	edgesLeft := int((float32(vertices) * sparseness)) - vertices
	for n := 0; n < edgesLeft; n++ {
		source := graph.RandomVertex()

		// Make sure the destination is not the source
		destination := source
		for source == destination {
			destination = graph.RandomVertex()
		}

		source.Connect(destination, rand.Intn(maxCost))
	}

	return graph
}

// Finds the cheapest path between two vertices in a graph
// Returns the destination vertex after it has been filled with the best path
func (graph *Graph) Dijkstra(origin *Vertex, destination *Vertex) *Vertex {
	// Add all the vertices to a set
	vertices := NewVertexSet()
	for _, vertex := range graph.Vertices {
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

	return destination
}

func main() {
	graph := GenerateGraph(7, 2.0, 10)
	origin := graph.FindVertex("A")
	destination := graph.FindVertex("G")

	graph.Display()

	fmt.Printf("Finding path between %v and %v...\n", origin.Id, destination.Id)
	graph.Dijkstra(origin, destination)

	fmt.Println("Result:")
	destination.DisplayPath()
}
