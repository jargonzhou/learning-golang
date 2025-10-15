package context_example

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//
// Context timeout
//

// timeout in child context is bound by timeout on the parent contxt
func TestTimeoutBound(t *testing.T) {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()

	start := time.Now()
	<-child.Done()
	end := time.Now()
	lastTime := end.Sub(start).Truncate(time.Second)
	fmt.Println("last", lastTime) // last 2s
}

// cancellation and timeout
// see: context_cancellation_test.go
//
// context.Cause(ctx)
// ctx.Err()
func TestContextCancellationAndTimeout(t *testing.T) {
	// WithTimeOut/WithDeadline, WithCancelCause
	// or
	// WithCancel, WithTimeoutCause/WithDeadlineCause
	ctx, cancelFuncParent := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFuncParent()
	// to return the error for cancellation
	ctx, cancelFunc := context.WithCancelCause(ctx)
	defer cancelFunc(nil)

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// spawn 2 goroutines
	spawnRandomStatus(ctx, cancelFuncParent, &wg, ch)
	spawnDelay(ctx, cancelFuncParent, &wg, ch)

loop:
	for {
		select {
		case s := <-ch:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			fmt.Println("in main: canceled with cause:", context.Cause(ctx), "err:", ctx.Err())
			break loop
		}
	}

	wg.Wait()
}
