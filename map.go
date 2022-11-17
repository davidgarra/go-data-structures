package main

import (
	"fmt"
)

type Bucket[K string, V any] struct {
	key   K
	value V
}

type Map[K string, V any] struct {
	Data [][]Bucket[K, V]
}

func NewMap[K string, V any](size int) Map[K, V] {
	return Map[K, V]{Data: make([][]Bucket[K, V], size)}
}

func (m *Map[K, V]) hash(key K) uint {
	var hash uint
	for _, b := range key {
		hash += uint(b)
		hash %= uint(len(m.Data))
	}

	return hash
}

func findBucket[K string, V any](bl []Bucket[K, V], key K) (Bucket[K, V], bool) {
	for _, b := range bl {
		if b.key == key {
			return b, true
		}
	}

	return Bucket[K, V]{}, false
}

func (m *Map[K, V]) Set(key K, value V) {
	hk := m.hash(key)
	bucketList := m.Data[hk]
	if _, found := findBucket(bucketList, key); found {
		fmt.Printf("Key %v already exists\n", key)
		return
	}
	bucketList = append(bucketList, Bucket[K, V]{key, value})
	m.Data[hk] = bucketList
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	hk := m.hash(key)
	bucket, found := findBucket(m.Data[hk], key)
	return bucket.value, found
}
