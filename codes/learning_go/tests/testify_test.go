// Tests with testify.
// more: https://github.com/stretchr/testify
package tests

import (
	"fmt"
	"testing"

	fib "learning_go/fibonacci"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	fmt.Printf("Testing A\n")
	require.Equal(t, 55, fib.FibonacciA(10))
}

func TestB(t *testing.T) {
	fmt.Printf("Testing B\n")
	require.Equal(t, 55, fib.FibonacciB(10))
}

func TestAB(t *testing.T) {
	fmt.Printf("Testing AB\n")
	require.Equal(t, 55, fib.FibonacciA(10))
	require.Equal(t, 55, fib.FibonacciB(10))
}

func TestPanic(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		panic("expected panic")
	}, "the code should have panic")
}
