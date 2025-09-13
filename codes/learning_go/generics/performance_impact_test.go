package generics

import "testing"

type Ager interface {
	age() int
}

func doubleAge(a Ager) int {
	return a.age() * 2
}

func doubleAgeGeneric[T Ager](a T) int {
	return a.age() * 2
}

type SAger struct {
	agei int
}

func (sa SAger) age() int {
	return sa.agei
}

var sa = SAger{agei: 42}

// 1000000000	         0.0000001 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDoubleAg(b *testing.B) {
	doubleAge(sa)
}

// 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDoubleAgGeneric(b *testing.B) {
	doubleAgeGeneric(sa)
}
