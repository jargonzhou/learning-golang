package fibonacci

// Slow
func FibonacciA(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciA(n-1) + FibonacciA(n-2)
}

// Fast
func FibonacciB(n int) int {
	if n <= 1 {
		return n
	}
	fib, fibMinus1 := 1, 0
	for ii := 2; ii <= n; ii++ {
		newFib := fib + fibMinus1
		fib, fibMinus1 = newFib, fib
	}
	return fib
}
