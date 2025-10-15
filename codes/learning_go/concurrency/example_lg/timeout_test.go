package example_lg

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// Example of timeout functions. - 'Learning Go': 12. Concurrency in Go, P.304
// use context.WithTimeout()

func timeLimit[T any](work func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()

	// is this goroutine leak???
	go func() {
		out <- work()
	}()

	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("time out")
	}
}

func doWork() int {
	time.Sleep(3 * time.Second) // 3s
	fmt.Println("return 42")
	return 42
}

func TestTimeoutFunction(t *testing.T) {
	result, err := timeLimit(doWork, 2*time.Second) // 2s
	fmt.Println(result, err)

	time.Sleep(5 * time.Second)
}
