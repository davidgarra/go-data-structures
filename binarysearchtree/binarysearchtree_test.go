package binarysearchtree

import (
	"datastructures/utils"
	"testing"
)

type Cint int

func (n Cint) Compare(other Cint) int {
	return int(n - other)
}

func checkLength(t *testing.T, bst *BinarySearchTree[Cint], expected int) {
	if length := bst.Length; length != expected {
		t.Fatalf("BinarySearchTree.Length = %v, want %v", length, expected)
	}
}

func checkLookup(t *testing.T, bst *BinarySearchTree[Cint], value Cint, expected Comparable[Cint], expectedFound bool) {
	if foundValue, found := bst.Lookup(value); found != expectedFound || foundValue != expected {
		t.Fatalf("BinarySearchTree.Lookup(%v) = (%v, %v), want (%v, %v)", foundValue, foundValue, found, expected, expectedFound)
	}
}

func checkRemove(t *testing.T, bst *BinarySearchTree[Cint], value Cint, expected bool) {
	if found := bst.Remove(value); found != expected {
		t.Fatalf("BinarySearchTree.Remove(%v) = %v, want %v", value, found, expected)
	}
}

func TestBinarySearchTreeInsert(t *testing.T) {
	utils.Quiet()
	bst := New[Cint]()

	checkLength(t, &bst, 0)
	bst.Insert(Cint(10))
	checkLength(t, &bst, 1)
	bst.Insert(Cint(1))
	checkLength(t, &bst, 2)
	bst.Insert(Cint(5))
	checkLength(t, &bst, 3)
	bst.Insert(Cint(7))
}

func TestBinarySearchTreeLookup(t *testing.T) {
	utils.Quiet()
	bst := New[Cint]()

	bst.Insert(Cint(10))
	bst.Insert(Cint(1))
	bst.Insert(Cint(5))
	bst.Insert(Cint(7))

	checkLookup(t, &bst, Cint(10), Cint(10), true)
	checkLookup(t, &bst, Cint(1), Cint(1), true)
	checkLookup(t, &bst, Cint(5), Cint(5), true)
	checkLookup(t, &bst, Cint(7), Cint(7), true)
	checkLookup(t, &bst, Cint(9), nil, false)
}

func TestBinarySearchTreeRemove(t *testing.T) {
	//utils.Quiet()
	bst := New[Cint]()
	
	checkRemove(t, &bst, Cint(10), false)
	checkLength(t, &bst, 0)
	bst.Insert(Cint(10))
	checkRemove(t, &bst, Cint(10), true)
	checkLength(t, &bst, 0)
	checkRemove(t, &bst, Cint(10), false)
	checkLength(t, &bst, 0)
	bst.Insert(Cint(1))
	bst.Insert(Cint(5))
	bst.Insert(Cint(7))
	checkRemove(t, &bst, Cint(10), false)
	checkLength(t, &bst, 3)
	checkRemove(t, &bst, Cint(5), true)
	checkLength(t, &bst, 2)
	checkRemove(t, &bst, Cint(1), true)
	checkLength(t, &bst, 1)
	checkRemove(t, &bst, Cint(7), true)
	checkLength(t, &bst, 0)
	checkRemove(t, &bst, Cint(7), false)
	checkLength(t, &bst, 0)
}
