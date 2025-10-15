package example_lg

import (
	"context"
	"fmt"
	"testing"
)

// Example of goroutines. - 'Learning Go': 12. Concurrency in Go, P.290
// main groutine launche multiple process groutines, communicate using channels.

const numGoroutines = 5

func process(val int) int {
	return val * 2
}

func processConcurrently(inVals []int) []int {
	// create channels: buffered
	in := make(chan int, numGoroutines)
	out := make(chan int, numGoroutines)

	// launch goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for val := range in {
				result := process(val)
				out <- result
			}
		}()
	}

	// load data into channel in another goroutine
	go func() {
		for _, v := range inVals {
			in <- v
		}
		close(in) // close in channel
	}()

	// read the data
	outVals := make([]int, 0, len(inVals))
	for i := 0; i < len(inVals); i++ {
		outVals = append(outVals, <-out) // read from out channel
	}
	return outVals
}

func TestProcessConcurrently(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := processConcurrently(x)
	fmt.Println(result)
}

// Example of terminating goroutines. - 'Learning Go': 12. Concurrency in Go, P.300

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			// on done
			case <-ctx.Done():
				return
			// on send data
			case ch <- i:
			}
		}
	}()
	return ch
}

func TestTerminateGoroutine(t *testing.T) {
	// use context to make sure goroutine terminated.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
