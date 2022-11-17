package main

type Bucket[K string, V any] struct {
	key   K
	value V
}

type Map[K string, V any] struct {
	Data []Bucket[K, V]
}

func NewMap[K string, V any](size int) Map[K, V] {
	return Map[K, V]{Data: make([]Bucket[K, V], size)}
}

func (m *Map[K, V]) hash(key K) uint {
	var hash uint
	for _, b := range key {
		hash += uint(b)
		hash %= uint(len(m.Data))
	}

	return hash
}

func (m *Map[K, V]) Set(key K, value V) {
	hk := m.hash(key)
	m.Data[hk] = Bucket[K, V]{key, value}
}

func (m *Map[K, V]) Get(key K) V {
	hk := m.hash(key)
	return m.Data[hk].value
}
