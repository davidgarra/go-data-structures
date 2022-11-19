package hashtable

import (
	"fmt"
)

type Bucket[V any] struct {
	key   string
	value V
}

type HashTable[V any] struct {
	data [][]Bucket[V]
}

func New[V any](size int) HashTable[V] {
	return HashTable[V]{data: make([][]Bucket[V], size)}
}

func (m *HashTable[V]) hash(key string) uint {
	var hash uint
	for _, b := range key {
		hash += uint(b)
		hash %= uint(len(m.data))
	}

	return hash
}

func findBucket[V any](bl []Bucket[V], key string) (Bucket[V], bool) {
	for _, b := range bl {
		if b.key == key {
			return b, true
		}
	}

	return Bucket[V]{}, false
}

func (m *HashTable[V]) Set(key string, value V) bool {
	hk := m.hash(key)
	bucketList := m.data[hk]
	if _, found := findBucket(bucketList, key); found {
		fmt.Printf("Key %v already exists\n", key)
		return false
	}
	bucketList = append(bucketList, Bucket[V]{key, value})
	m.data[hk] = bucketList
	return true
}

func (m *HashTable[V]) Get(key string) (V, bool) {
	hk := m.hash(key)
	bucket, found := findBucket(m.data[hk], key)
	return bucket.value, found
}

func (m *HashTable[V]) Keys() []string {
	var keys []string
	for _, bucketlist := range m.data {
		for _, bucket := range bucketlist {
			keys = append(keys, bucket.key)
		}
	}
	return keys
}
