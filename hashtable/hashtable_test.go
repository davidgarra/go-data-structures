package hashtable

import (
	"datastructures/utils"
	"testing"
)

func TestHashTableSet(t *testing.T) {
	utils.Quiet()
	myHashTable := New[string](2)

	trySet := func(key string, value string, expected bool) {
		inserted := myHashTable.Set(key, value)
		if inserted != expected {
			t.Fatalf("Insert '%v:%v' in hashtable = %v, want %v", key, value, inserted, expected)
		}
	}
	trySet("k1", "v1", true)  // test first insert
	trySet("k1", "v1", false) // test already present key
	trySet("k2", "v2", true)  // test second insert
	trySet("k3", "v3", true)  // test collision
}

func TestHashTableGet(t *testing.T) {
	utils.Quiet()
	myHashTable := New[string](2)

	tryGet := func(key string, value string, expected bool) {
		if got, found := myHashTable.Get(key); found != expected || got != value {
			t.Fatalf("myHashTable.Get(%v)=(%v:%v), want (%v, %v)", key, got, found, value, expected)
		}
	}
	myHashTable.Set("k1", "v1")
	myHashTable.Set("k2", "v2")
	myHashTable.Set("k2", "v2")
	myHashTable.Set("k3", "v3")

	tryGet("k1", "v1", true) // test first get
	tryGet("k1", "v1", true) // test second get on the same element
	tryGet("k2", "v2", true) // test get on double inserted element
	tryGet("k3", "v3", true) // test get on collided element
	tryGet("k4", "", false)  // test get on non existing element
}

func TestHashTableKeys(t *testing.T) {
	utils.Quiet()
	myHashTable := New[string](2)

	checkKey := func(key string, keys []string, expected bool) {
		for _, a := range keys {
			if a == key {
				if expected == false {
					t.Fatalf("False positive in myHashTable.Keys(), keys: %v, false positive: %v", keys, key)
				}
				return
			}
		}

		if expected == true {
			t.Fatalf("Key missing in myHashTable.Keys(), keys: %v, want: %v", keys, key)
		}
	}

	myHashTable.Set("k1", "v1")
	myHashTable.Set("k2", "v2")
	myHashTable.Set("k2", "v2")
	myHashTable.Set("k3", "v3")

	keys := myHashTable.Keys()
	if len := len(keys); len != 3 {
		t.Fatalf("Keys lenght=%v, want %v", len, 3)
	}

	checkKey("k1", keys, true)  // test first inserted key
	checkKey("k2", keys, true)  // test second inserted key
	checkKey("k3", keys, true)  // test collided key
	checkKey("k4", keys, false) // test missing key
}
