package example_lg

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// Example of backpressure. - 'Learning Go': 12. Concurrency in Go, P.303
type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	return &PressureGauge{
		// buffered channel
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(ctx context.Context, f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	case <-ctx.Done():
		fmt.Println("process canceled")
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func restrictedResource() string {
	time.Sleep(100 * time.Millisecond)
	return "done"
}

func TestBackPressure(t *testing.T) {
	pg := New(10)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// mock concurrent accesses to restricted resources.
	for i := 0; i < 20; i++ {
		go func(ctx context.Context) {
			err := pg.Process(ctx, func() {
				fmt.Println(restrictedResource())
			})
			if err != nil {
				fmt.Println(i, "too many requests")
			} else {
				select {
				case <-ctx.Done():
					fmt.Println(i, "goroutine canceled")
				default:
					fmt.Println(i, "goroutine done")
				}
			}
		}(ctx)
	}

	<-ctx.Done()
}
