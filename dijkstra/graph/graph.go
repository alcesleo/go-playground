package graph

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
	fmt.Println("Graph:")
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
	result := []*Vertex{v} // Start from here

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
