package graph

import (
	"datastructures/utils"
	"testing"
)

func checkLength(t *testing.T, g *Graph[int], expected int) {
	if length := g.Length; length != expected {
		t.Fatalf("Graph.Length = %v, want %v", length, expected)
	}
}

func checkGet(t *testing.T, g *Graph[int], key string, expected int, expectedFound bool) {
	if value, found := g.Get(key); found != expectedFound || value != expected {
		t.Fatalf("Graph.Get(%v) = (%v, %v), want (%v, %v)", key, value, found, expected, expectedFound)
	}
}

func checkRemoveVertex(t *testing.T, g *Graph[int], key string, expected bool) {
	if found := g.RemoveVertex(key); found != expected {
		t.Fatalf("Graph.RemoveVertex(%v) = %v, want %v", key, found, expected)
	}
}

func checkCheckEdge(t *testing.T, g *Graph[int], first string, second string, expected bool) {
	if found := g.CheckEdge(first, second); found != expected {
		t.Fatalf("Graph.CheckEdge(%v, %v) = %v, want %v", first, second, found, expected)
	}
}

func TestGraphAddVertex(t *testing.T) {
	utils.Quiet()
	g := New[int]()

	checkLength(t, &g, 0)
	g.AddVertex("A", 1)
	checkLength(t, &g, 1)
	g.AddVertex("B", 2)
	checkLength(t, &g, 2)
}

func TestGraphRemoveVertex(t *testing.T) {
	utils.Quiet()
	g := New[int]()

	checkRemoveVertex(t, &g, "A", false)
	checkLength(t, &g, 0)
	g.AddVertex("A", 1)
	g.AddVertex("B", 1)
	g.AddVertex("C", 1)


	g.AddEdge("A", "B")
	g.AddEdge("C", "B")

	checkRemoveVertex(t, &g, "A", true)
	checkLength(t, &g, 2)
	checkRemoveVertex(t, &g, "B", true)
	checkLength(t, &g, 1)

	checkCheckEdge(t, &g, "A", "B", false)
	checkCheckEdge(t, &g, "B", "A", false)
	checkCheckEdge(t, &g, "B", "C", false)
	checkCheckEdge(t, &g, "C", "B", false)

	checkRemoveVertex(t, &g, "C", true)
	checkLength(t, &g, 0)
}

func TestGraphGet(t *testing.T) {
	utils.Quiet()
	g := New[int]()

	checkGet(t, &g, "A", 0, false)
	g.AddVertex("A", 1)
	checkGet(t, &g, "A", 1, true)
	g.AddVertex("B", 2)
	checkGet(t, &g, "B", 2, true)
	g.AddVertex("C", 3)
	checkGet(t, &g, "C", 3, true)
	checkGet(t, &g, "D", 0, false)
}

func TestGraphAddEdge(t *testing.T) {
	utils.Quiet()
	g := New[int]()

	checkLength(t, &g, 0)
	g.AddVertex("A", 1)
	g.AddVertex("B", 1)
	g.AddVertex("C", 1)

	checkCheckEdge(t, &g, "A", "B", false)
	checkCheckEdge(t, &g, "B", "A", false)
	g.AddEdge("A", "B")
	g.AddEdge("C", "B")
	checkCheckEdge(t, &g, "A", "B", true)
	checkCheckEdge(t, &g, "B", "A", true)

	checkCheckEdge(t, &g, "A", "C", false)
	checkCheckEdge(t, &g, "C", "A", false)

	checkCheckEdge(t, &g, "C", "B", true)
	checkCheckEdge(t, &g, "B", "C", true)

	checkCheckEdge(t, &g, "A", "D", false)
	checkCheckEdge(t, &g, "D", "A", false)
}

func TestGraphRemoveEdge(t *testing.T) {
	utils.Quiet()
	g := New[int]()

	checkLength(t, &g, 0)
	g.AddVertex("A", 1)
	g.AddVertex("B", 1)
	g.AddVertex("C", 1)

	g.AddEdge("A", "B")
	g.AddEdge("C", "B")
	checkCheckEdge(t, &g, "A", "B", true)
	checkCheckEdge(t, &g, "B", "A", true)
	checkCheckEdge(t, &g, "A", "C", false)
	checkCheckEdge(t, &g, "C", "A", false)

	g.RemoveEdge("A", "B")
	checkCheckEdge(t, &g, "A", "B", false)
	checkCheckEdge(t, &g, "B", "A", false)
	checkCheckEdge(t, &g, "A", "C", false)
	checkCheckEdge(t, &g, "C", "A", false)
}