// Example Benchmarks.
package tests

import (
	fib "learning_go/fibonacci"
	"testing"
)

func BenchmarkFibonacciA32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.FibonacciA(32)
	}
}

func BenchmarkFibonacciB32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.FibonacciB(32)
	}
}
