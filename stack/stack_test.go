package stack

import (
	"datastructures/utils"
	"testing"
)

func checkLength(t *testing.T, s *Stack[string], expected int) {
	if length := s.length; length != expected {
		t.Fatalf("Stack.length = %v, want %v", length, expected)
	}
}

func checkPeek(t *testing.T, s *Stack[string], expected string, expectedFound bool) {
	if item, found := s.Peek(); item != expected || found != expectedFound {
		t.Fatalf("Stack.Peek() = %v, want %v", item, expected)
	}
}

func checkPop(t *testing.T, s *Stack[string], expected string, expectedFound bool) {
	if item, found := s.Pop(); item != expected || found != expectedFound {
		t.Fatalf("Stack.Pop() = %v, want %v", item, expected)
	}
}

func checkIsEmpty(t *testing.T, s *Stack[string], expected bool) {
	if isEmpty := s.IsEmpty(); isEmpty != expected {
		t.Fatalf("Stack.IsEmpty() = %v, want %v", isEmpty, expected)
	}
}

func TestStackPush(t *testing.T) {
	utils.Quiet()
	s := New[string]()

	checkLength(t, &s, 0)
	s.Push("A")
	checkLength(t, &s, 1)
	s.Push("B")
	checkLength(t, &s, 2)
}

func TestStackPeek(t *testing.T) {
	utils.Quiet()
	s := New[string]()

	checkPeek(t, &s, "", false)
	s.Push("A")
	checkPeek(t, &s, "A", true)
	s.Push("B")
	checkPeek(t, &s, "B", true)
}

func TestStackPop(t *testing.T) {
	utils.Quiet()
	s := New[string]()

	checkPop(t, &s, "", false)
	s.Push("A")
	s.Push("B")
	checkPop(t, &s, "B", true)
	checkPop(t, &s, "A", true)
	checkPop(t, &s, "", false)
}

func TestStackIsEmpty(t *testing.T) {
	utils.Quiet()
	s := New[string]()

	checkIsEmpty(t, &s, true)
	s.Push("A")
	checkIsEmpty(t, &s, false)
	s.Peek()
	checkIsEmpty(t, &s, false)
	s.Pop()
	checkIsEmpty(t, &s, true)
}
