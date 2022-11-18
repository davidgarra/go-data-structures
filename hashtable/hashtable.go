package hashtable

import (
	"fmt"
)

// TODO: find a way to use a key different from string
type Bucket[K string, V any] struct {
	key   K
	value V
}

type HashTable[K string, V any] struct {
	Data [][]Bucket[K, V]
}

func New[K string, V any](size int) HashTable[K, V] {
	return HashTable[K, V]{Data: make([][]Bucket[K, V], size)}
}

func (m *HashTable[K, V]) hash(key K) uint {
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

func (m *HashTable[K, V]) Set(key K, value V) bool {
	hk := m.hash(key)
	bucketList := m.Data[hk]
	if _, found := findBucket(bucketList, key); found {
		fmt.Printf("Key %v already exists\n", key)
		return false
	}
	bucketList = append(bucketList, Bucket[K, V]{key, value})
	m.Data[hk] = bucketList
	return true
}

func (m *HashTable[K, V]) Get(key K) (V, bool) {
	hk := m.hash(key)
	bucket, found := findBucket(m.Data[hk], key)
	return bucket.value, found
}

func (m *HashTable[K, V]) Keys() []K {
	var keys []K
	for _, bucketlist := range m.Data {
		for _, bucket := range bucketlist {
			keys = append(keys, bucket.key)
		}
	}
	return keys
}
