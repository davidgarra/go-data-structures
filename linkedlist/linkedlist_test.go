package linkedlist

import (
	"datastructures/utils"
	"testing"
)

func checkLength(t *testing.T, l *LinkedList[string], expected int) {
	if length := l.Length; length != expected {
		t.Fatalf("LinkedList.Length = %v, want %v", length, expected)
	}
}

func checkGet(t *testing.T, l *LinkedList[string], index int, expected string, expectedFound bool) {
	if value, found := l.Get(index); found != expectedFound || value != expected {
		t.Fatalf("LinkedList.Get(%v) = (%v, %v), want (%v, %v)", index, value, found, expected, expectedFound)
	}
}

func checkRemove(t *testing.T, l *LinkedList[string], index int, expected bool) {
	if found := l.Remove(index); found != expected {
		t.Fatalf("LinkedList.Remove(%v) = %v, want %v", index, found, expected)
	}
}

func TestLinkedListAppend(t *testing.T) {

	utils.Quiet()
	l := New[string]()

	checkLength(t, &l, 0)
	l.Append("A")
	checkLength(t, &l, 1)
	checkGet(t, &l, 0, "A", true)
	l.Append("B")
	checkLength(t, &l, 2)
	checkGet(t, &l, 1, "B", true)
	l.Append("C")
	checkLength(t, &l, 3)
	checkGet(t, &l, 2, "C", true)
}

func TestLinkedListPrepend(t *testing.T) {

	utils.Quiet()
	l := New[string]()

	checkLength(t, &l, 0)
	l.Prepend("A")
	checkLength(t, &l, 1)
	checkGet(t, &l, 0, "A", true)
	l.Prepend("B")
	checkLength(t, &l, 2)
	checkGet(t, &l, 0, "B", true)
	l.Prepend("C")
	checkLength(t, &l, 3)
	checkGet(t, &l, 0, "C", true)
}

func TestLinkedListInsert(t *testing.T) {

	utils.Quiet()
	l := New[string]()

	checkLength(t, &l, 0)
	l.Insert("A", 0) // test first insert
	checkLength(t, &l, 1)
	checkGet(t, &l, 0, "A", true)
	l.Insert("B", 0) // test insert to head
	checkLength(t, &l, 2)
	checkGet(t, &l, 0, "B", true)
	l.Insert("C", -1) // test negative insert
	checkLength(t, &l, 3)
	checkGet(t, &l, 0, "C", true)
	l.Insert("D", 100) // test overflow insert
	checkLength(t, &l, 4)
	checkGet(t, &l, 3, "D", true)
	l.Insert("E", 2) // test middle insert
	checkLength(t, &l, 5)
	checkGet(t, &l, 2, "E", true)
	l.Insert("F", 5) // test tail node insert
	checkLength(t, &l, 6)
	checkGet(t, &l, 5, "F", true)
}

func TestLinkedListGet(t *testing.T) {

	utils.Quiet()
	l := New[string]()

	l.Append("A")
	l.Append("B")
	l.Append("C")

	checkGet(t, &l, 0, "A", true)  // test head node
	checkGet(t, &l, 1, "B", true)  // test middle node
	checkGet(t, &l, 2, "C", true)  // test tail node
	checkGet(t, &l, 3, "", false)  // test overflow
	checkGet(t, &l, -1, "", false) // test negative node
}

func TestLinkedListRemove(t *testing.T) {
	utils.Quiet()
	l := New[string]()

	checkRemove(t, &l, 0, false) // test empy list
	checkLength(t, &l, 0)
	l.Append("A")
	l.Append("B")
	l.Append("C")
	l.Append("D")
	checkRemove(t, &l, 10, false) // test missing node
	checkLength(t, &l, 4)
	checkRemove(t, &l, -10, false) // test negative node
	checkLength(t, &l, 4)
	checkRemove(t, &l, 0, true) // test head node
	checkLength(t, &l, 3)
	checkRemove(t, &l, 1, true) // test middle node
	checkLength(t, &l, 2)
	checkRemove(t, &l, 1, true) // test tail node
	checkLength(t, &l, 1)
	checkRemove(t, &l, 0, true) // test single node
	checkLength(t, &l, 0)
}
