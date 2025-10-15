package example_lg

import (
	"fmt"
	"sync"
	"testing"
)

// Example of sync.WaitGroup. - 'Learning Go': 12. Concurrency in Go, P.307

// T: input
// R: result
// num: number of goroutines
func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)

	// spawn num goroutines to process input
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()

			for v := range in {
				out <- processor(v)
			}
		}()
	}

	// wait for job done
	go func() {
		wg.Wait()
		close(out) // close output channel
	}()

	// gather results
	var result []R
	for v := range out {
		result = append(result, v)
	}
	return result
}

func TestProcessAndGather(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch) // close input channel
	}()

	results := processAndGather(ch, func(i int) int {
		return i * 2
	}, 3)
	fmt.Println(results)
}
