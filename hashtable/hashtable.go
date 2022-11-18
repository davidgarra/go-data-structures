package hashtable

import (
	"fmt"
)

type Bucket[V any] struct {
	key   string
	value V
}

type HashTable[V any] struct {
	Data [][]Bucket[V]
}

func New[V any](size int) HashTable[V] {
	return HashTable[V]{Data: make([][]Bucket[V], size)}
}

func (m *HashTable[V]) hash(key string) uint {
	var hash uint
	for _, b := range key {
		hash += uint(b)
		hash %= uint(len(m.Data))
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
	bucketList := m.Data[hk]
	if _, found := findBucket(bucketList, key); found {
		fmt.Printf("Key %v already exists\n", key)
		return false
	}
	bucketList = append(bucketList, Bucket[V]{key, value})
	m.Data[hk] = bucketList
	return true
}

func (m *HashTable[V]) Get(key string) (V, bool) {
	hk := m.hash(key)
	bucket, found := findBucket(m.Data[hk], key)
	return bucket.value, found
}

func (m *HashTable[V]) Keys() []string {
	var keys []string
	for _, bucketlist := range m.Data {
		for _, bucket := range bucketlist {
			keys = append(keys, bucket.key)
		}
	}
	return keys
}
