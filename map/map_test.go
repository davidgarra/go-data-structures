package _map

import (
	"log"
	"os"
	"testing"
)

func TestMapSet(t *testing.T) {
	quiet()
	myMap := New[string, string](2)

	trySet := func(key string, value string, expected bool) {
		inserted := myMap.Set(key, value)
		if inserted != expected {
			t.Fatalf("Insert '%v:%v' in map = %v, want %v", key, value, inserted, expected)
		}
	}
	trySet("k1", "v1", true)  // test first insert
	trySet("k1", "v1", false) // test already present key
	trySet("k2", "v2", true)  // test second insert
	trySet("k3", "v3", true)  // test collision
}

func TestMapGet(t *testing.T) {
	quiet()
	myMap := New[string, string](2)

	tryGet := func(key string, value string, expected bool) {
		if got, found := myMap.Get(key); found != expected || got != value {
			t.Fatalf("myMap.Get(%v)=(%v:%v), want (%v, %v)", key, got, found, value, expected)
		}
	}
	myMap.Set("k1", "v1")
	myMap.Set("k2", "v2")
	myMap.Set("k2", "v2")
	myMap.Set("k3", "v3")

	tryGet("k1", "v1", true) // test first get
	tryGet("k1", "v1", true) // test second get on the same element
	tryGet("k2", "v2", true) // test get on double inserted element
	tryGet("k3", "v3", true) // test get on collided element
	tryGet("k4", "", false)  // test get on non existing element
}

func TestMapKeys(t *testing.T) {
	quiet()
	myMap := New[string, string](2)

	checkKey := func(key string, keys []string, expected bool) {
		for _, a := range keys {
			if a == key {
				if expected == false {
					t.Fatalf("False positive in myMap.Keys(), keys: %v, false positive: %v", keys, key)
				}
				return
			}
		}

		if expected == true {
			t.Fatalf("Key missing in myMap.Keys(), keys: %v, want: %v", keys, key)
		}
	}

	myMap.Set("k1", "v1")
	myMap.Set("k2", "v2")
	myMap.Set("k2", "v2")
	myMap.Set("k3", "v3")

	keys := myMap.Keys()
	if len := len(keys); len != 3 {
		t.Fatalf("Keys lenght=%v, want %v", len, 3)
	}

	checkKey("k1", keys, true)  // test first inserted key
	checkKey("k2", keys, true)  // test second inserted key
	checkKey("k3", keys, true)  // test collided key
	checkKey("k4", keys, false) // test missing key
}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}
