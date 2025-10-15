package example_cig

import (
	"fmt"
	"testing"
	"time"
)

func TestOrDoneChannels(t *testing.T) {
	sig := func(after time.Duration) <-chan any {
		c := make(chan any)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-OrDoneChannels(
		sig(time.Minute),
		sig(2*time.Minute),
		sig(3*time.Minute),
		sig(time.Second),
		sig(time.Hour))
	fmt.Printf("done after %v\n", time.Since(start))
}
