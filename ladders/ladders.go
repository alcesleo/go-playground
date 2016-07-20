package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	limit := 5
	words := limitNumLetters(dictionary(), limit)
	graph := buildGraph(words)

	printShortestPath(graph, "cold", "warm")
	printShortestPath(graph, "chaos", "order")
	printShortestPath(graph, "right", "wrong")
}

func dictionary() []string {
	text, _ := ioutil.ReadFile("/usr/share/dict/words")
	return strings.Split(string(text), "\n")
}

func printShortestPath(graph map[string][]string, start string, end string) {
	path := findShortestPath(graph, start, end)
	fmt.Printf("%v steps between %v and %v\n", len(path)-1, start, end)
	fmt.Println(strings.Join(path, " -> "))
	fmt.Println()
}

func findShortestPath(graph map[string][]string, start string, end string) []string {
	dist := make(map[string]int)
	prev := make(map[string]string)
	queue := make([]string, 0)

	// Mark the distance to all nodes as infinite
	for word := range graph {
		dist[word] = math.MaxInt32
	}

	// Mark the start as free and enqueue it
	dist[start] = 0
	queue = append(queue, start)

	// Traverse the graph with a breadth first search
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, edge := range graph[current] {
			if dist[edge] == math.MaxInt32 {
				dist[edge] = dist[current] + 1
				prev[edge] = current
				queue = append(queue, edge)
			}
		}
	}

	// Walk backwards to extract the path from start to end
	result := []string{end}
	current := end

	for prev[current] != "" {
		current = prev[current]
		result = append([]string{current}, result...)
	}

	return result
}

// Represents the graph as a map of nodes to a list of their edges
func buildGraph(words []string) map[string][]string {
	// Put words in buckets
	buckets := make(map[string][]string)

	for _, word := range words {
		for char := range word {
			// Put the word "example" in the bucket "ex_mple"
			bucket := blankOutChar(word, char)
			buckets[bucket] = append(buckets[bucket], word)
		}
	}

	// Build a graph out of the buckets
	graph := make(map[string][]string)

	for _, edges := range buckets {
		for _, edge1 := range edges {
			for _, edge2 := range edges {
				if edge1 != edge2 {
					graph[edge1] = append(graph[edge1], edge2)
				}
			}
		}
	}

	return graph
}

func limitNumLetters(words []string, n int) []string {
	result := make([]string, 0)
	for _, word := range words {
		if len(word) <= n {
			result = append(result, word)
		}
	}
	return result
}

func blankOutChar(str string, n int) string {
	return str[:n] + "_" + str[n+1:]
}
