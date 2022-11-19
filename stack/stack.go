package stack

import (
	"datastructures/utils"
)

type Stack[T any] struct {
	data   []T
	length int
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
	s.length++
}

func (s *Stack[T]) Peek() (T, bool) {
	if len := len(s.data); len > 0 {
		return s.data[len-1], true
	}

	return utils.Zero[T](), false
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.length > 0 {
		last := s.data[s.length-1]
		s.data = s.data[:s.length-1]
		s.length--
		return last, true
	}

	return utils.Zero[T](), false
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}
