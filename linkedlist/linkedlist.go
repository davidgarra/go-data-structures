package linkedlist

import (
	"datastructures/utils"
	"fmt"
	"strings"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (l *LinkedList[T]) insertFirst(value T) {
	newNode := Node[T]{value, nil}
	l.head = &newNode
	l.tail = &newNode
	l.length = 1
}

func (l *LinkedList[T]) Append(value T) {
	if l.length == 0 {
		l.insertFirst(value)
		return
	}
	newNode := Node[T]{value, nil}
	l.tail.next = &newNode
	l.tail = &newNode
	l.length++
}

func (l *LinkedList[T]) Prepend(value T) {
	if l.length == 0 {
		l.insertFirst(value)
		return
	}
	newNode := Node[T]{value, nil}
	newNode.next = l.head
	l.head = &newNode
	l.length++
}

func (l *LinkedList[T]) getNode(index int) (*Node[T], bool) {
	if index < 0 || index >= l.length {
		return nil, false
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, true
}

func (l *LinkedList[T]) Get(index int) (T, bool) {
	node, found := l.getNode(index)
	if !found {
		return utils.Zero[T](), false
	}
	return node.value, found
}

func (l *LinkedList[T]) Insert(value T, index int) {
	if l.length == 0 {
		l.insertFirst(value)
		return
	}

	if index >= l.length {
		l.Append(value)
		return
	}
	if index <= 0 {
		l.Prepend(value)
		return
	}

	newNode := Node[T]{value, nil}
	previous, _ := l.getNode(index - 1)
	newNode.next = previous.next
	previous.next = &newNode
	l.length++
}

func (l *LinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= l.length {
		return false
	}
	if index == 0 {
		l.head = l.head.next
	} else {
		previous, _ := l.getNode(index - 1)
		unwanted := previous.next
		previous.next = unwanted.next
		if index == l.length {
			l.tail = previous
		}
	}
	l.length--
	return true
}

func (l *LinkedList[T]) toArray() []T {
	array := make([]T, l.length)
	current := l.head
	for i := 0; i < l.length; i++ {
		array[i] = current.value
		current = current.next
	}
	return array
}

func (l *LinkedList[T]) String() string {
	strArr := make([]string, l.length)
	for i, val := range l.toArray() {
		strArr[i] = fmt.Sprintf("%v", val)
	}
	return strings.Join(strArr, " -> ")
}
