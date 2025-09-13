package generics

import (
	"fmt"
	"testing"
)

// type paramter: T
// type constraint: any
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func TestStack(t *testing.T) {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	fmt.Printf("%v\n", intStack)
	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(5))

	// ERROR: cannot use "nope" (untyped string constant) as int value in argument to intStack.Push
	// var s Stack[int]
	// s.Push("nope")
}
