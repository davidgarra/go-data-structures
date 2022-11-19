package queue

import "testing"

func checkLength(t *testing.T, q *Queue[string], expected int) {
	if length := q.Length; length != expected {
		t.Fatalf("Queue.Length = %v, want %v", length, expected)
	}
}

func checkDequeue(t *testing.T, q *Queue[string], expected string, expectedFound bool) {
	if item, found := q.Dequeue(); item != expected || found != expectedFound {
		t.Fatalf("Queue.Dequeue() = %v, want %v", item, expected)
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := New[string]()

	checkLength(t, &q, 0)
	q.Enqueue("A")
	checkLength(t, &q, 1)
	q.Enqueue("B")
	checkLength(t, &q, 2)
}

func TestQueueDequeue(t *testing.T) {
	q := New[string]()

	checkDequeue(t, &q, "", false)
	q.Enqueue("A")
	checkDequeue(t, &q, "A", true)
	q.Enqueue("B")
	q.Enqueue("C")
	checkDequeue(t, &q, "B", true)
	checkDequeue(t, &q, "C", true)
	checkDequeue(t, &q, "", false)
}
