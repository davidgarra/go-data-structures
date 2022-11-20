package binarysearchtree

import (
	"datastructures/utils"
	"fmt"
)

type Comparable[T any] interface {
	Compare(T) int
}

type node[T any] struct {
	value Comparable[T]
	left  *node[T]
	right *node[T]
}

type BinarySearchTree[T any] struct {
	root   *node[T]
	Length int
}

func New[T any]() BinarySearchTree[T] {
	return BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Insert(value Comparable[T]) bool {
	newNode := node[T]{value, nil, nil}
	if bst.root == nil {
		bst.root = &newNode
		bst.Length++
		return true
	}

	currentNode := bst.root
	for {
		comparison := value.Compare(currentNode.value.(T))
		if comparison > 0 {
			if currentNode.right == nil {
				currentNode.right = &newNode
				bst.Length++
				return true
			}
			currentNode = currentNode.right

		} else if comparison < 0 {
			if currentNode.left == nil {
				currentNode.left = &newNode
				bst.Length++
				return true
			}
			currentNode = currentNode.left
		} else {
			fmt.Printf("Value %v already exists in the binary tree", value)
			return false
		}
	}
}

func (bst *BinarySearchTree[T]) Lookup(value Comparable[T]) (Comparable[T], bool) {
	currentNode := bst.root
	for {
		if currentNode == nil {
			return utils.Zero[Comparable[T]](), false
		}
		comparison := value.Compare(currentNode.value.(T))
		if comparison > 0 {
			currentNode = currentNode.right
		} else if comparison < 0 {
			currentNode = currentNode.left
		} else {
			return currentNode.value, true
		}
	}
}

func (bst *BinarySearchTree[T]) Remove(value Comparable[T]) bool {
	currentNode := bst.root
	var parentNode *node[T]
	for {
		if currentNode == nil {
			break
		}
		comparison := value.Compare(currentNode.value.(T))
		if comparison > 0 {
			parentNode = currentNode
			currentNode = currentNode.right
		} else if comparison < 0 {
			parentNode = currentNode
			currentNode = currentNode.left
		} else {
			break
		}
	}

	if currentNode == nil {
		return false
	} else {
		if currentNode.right == nil {
			// no right child
			if parentNode == nil {
				bst.root = currentNode.left
			} else {
				if currentNode.value.Compare(parentNode.value.(T)) < 0 {
					// parent > current value
					parentNode.left = currentNode.left
				} else if currentNode.value.Compare(parentNode.value.(T)) > 0 {
					// parent < current value
					parentNode.right = currentNode.left
				}
			}
		} else if currentNode.right.left == nil {
			// right child with no left child
			currentNode.right.left = currentNode.left
			if parentNode == nil {
				bst.root = currentNode.right
			} else {
				if currentNode.value.Compare(parentNode.value.(T)) < 0 {
					// parent > current value
					parentNode.left = currentNode.right
				} else if currentNode.value.Compare(parentNode.value.(T)) > 0 {
					parentNode.right = currentNode.right
				}
			}
		} else {
			// right child with left child
			innerleft := currentNode.right.left
			innerleftParent := currentNode.right
			for {
				if innerleft.left == nil {
					break
				}
				innerleftParent = innerleft
				innerleft = innerleft.left
			}

			innerleftParent.left = innerleft.right
			innerleft.left = currentNode.left
			innerleft.right = currentNode.right

			if parentNode == nil {
				bst.root = innerleft
			} else {
				if currentNode.value.Compare(parentNode.value.(T)) < 0 {
					parentNode.left = innerleft
				} else if currentNode.value.Compare(parentNode.value.(T)) > 0 {
					parentNode.right = innerleft
				}
			}
		}
	}
	bst.Length--
	return true
}
