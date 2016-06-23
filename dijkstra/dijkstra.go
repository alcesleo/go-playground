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

func (g *Graph) Display() {
	for _, vertex := range g.Vertices {
		fmt.Printf(" - %v\n", vertex.Id)
		for _, edge := range vertex.Edges {
			fmt.Printf("   |-> %v (%v)\n", edge.Destination.Id, edge.Weight)
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
	result := []*Vertex{v} // Start from the here

	for target.Prev != nil {
		result = append([]*Vertex{target.Prev}, result...)
		target = target.Prev
	}
	return result
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

// Generate data

func GenerateGraph() *Graph {
	graph := NewGraph()

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

// Finds the cheapest path between two vertices in a graph
// Returns a slice of the steps and the total cost of those steps
func (graph *Graph) Dijkstra(origin *Vertex, destination *Vertex) ([]*Vertex, int) {
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

	return destination.Path(), destination.Cost
}

func main() {
	graph := GenerateGraph()

	fmt.Println("Graph:")
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
	fmt.Printf(" (%v)\n", cost)
}
