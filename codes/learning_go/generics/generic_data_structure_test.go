package generics

import (
	"cmp"
	"fmt"
	"testing"
)

// generic function
type OrderableFunc[T any] func(t1, t2 T) int

// generic struct
// WARN: rename Node to TNode, Tree to TTree due to previous definitions
type TNode[T any] struct {
	val         T
	left, right *TNode[T]
}

type TTree[T any] struct {
	f    OrderableFunc[T]
	root *TNode[T]
}

// consturct a new tree
func NewTree[T any](f OrderableFunc[T]) *TTree[T] {
	return &TTree[T]{
		f: f,
	}
}

func (t *TTree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *TTree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *TNode[T]) Add(f OrderableFunc[T], v T) *TNode[T] {
	if n == nil {
		return &TNode[T]{val: v}
	}

	switch r := f(v, n.val); {
	case r <= -1: // new value in less than current value
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *TNode[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}

	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

func TestGenericDataStructures(t *testing.T) {
	t1 := NewTree(cmp.Compare[int])
	fmt.Printf("%v\n", t1)
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)

	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))
}
