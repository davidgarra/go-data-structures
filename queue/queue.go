package queue

import (
	"datastructures/linkedlist"
)

type Queue[T any] struct {
	data   linkedlist.LinkedList[T]
	Length int
}

func New[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.data.Prepend(item)
	q.Length++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	index := q.Length
	item, found := q.data.Get(index)
	q.data.Remove(index)
	q.Length--
	return item, found
}

func (q *Queue[T]) String() string {
	return q.data.String()
}
