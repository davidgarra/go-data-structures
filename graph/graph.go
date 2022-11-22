package graph

import "datastructures/utils"

type vertex[T any] struct {
	key       string
	value     T
	neighbors map[string]*vertex[T]
}

type Graph[T any] struct {
	data   map[string]*vertex[T]
	Length int
}

func New[T any]() Graph[T] {
	return Graph[T]{data: make(map[string]*vertex[T])}
}

func (g *Graph[T]) AddVertex(key string, value T) bool {
	newVertex := vertex[T]{key, value, map[string]*vertex[T]{}}
	if g.data[key] != nil {
		return false
	}
	g.data[key] = &newVertex
	g.Length++
	return true
}

func (g *Graph[T]) AddEdge(first string, second string) bool {
	firstVertex := g.data[first]
	secondVertex := g.data[second]
	if firstVertex == nil || secondVertex == nil {
		return false
	}
	firstVertex.neighbors[second] = secondVertex
	secondVertex.neighbors[first] = firstVertex
	return true
}

func (g *Graph[T]) RemoveEdge(first string, second string) bool {
	if !g.CheckEdge(first, second) {
		return false
	}
	firstVertex := g.data[first]
	secondVertex := g.data[second]
	delete(firstVertex.neighbors, second)
	delete(secondVertex.neighbors, first)
	return true
}

func (g *Graph[T]) RemoveVertex(key string) bool {
	vertex := g.data[key]
	if vertex == nil {
		return false
	}
	for _, neighbor := range vertex.neighbors {
		delete(neighbor.neighbors, key)
	}
	delete(g.data, key)
	g.Length--
	return true
}

func (g *Graph[T]) Get(key string) (T, bool) {
	vertex := g.data[key]
	if vertex == nil {
		return utils.Zero[T](), false
	}
	return vertex.value, true
}

func (g *Graph[T]) CheckEdge(first string, second string) bool {
	firstVertex := g.data[first]
	secondVertex := g.data[second]
	if firstVertex == nil || secondVertex == nil {
		return false
	}
	if firstVertex.neighbors[secondVertex.key] == nil {
		return false
	}
	return true
}

func (g *Graph[T]) String() string {
	str := ""
	for key, vertex := range g.data {
		str += "(" + key + ")"
		length := len(vertex.neighbors)
		if length == 0 {
			str += "\n"
			continue
		}
		str += " --> { "
		i := 0
		for _, neighbor := range vertex.neighbors {
			i++
			str += "(" + neighbor.key + ")"
			if i < length {
				str += ", "
			}
		}
		str += " }\n"
	}
	return str
}
