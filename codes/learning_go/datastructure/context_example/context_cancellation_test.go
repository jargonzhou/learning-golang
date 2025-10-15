package context_example

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"testing"
	"time"
)

//
// Context cancellation
//

// HTTP client
var client = &http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout: 3 * time.Second,
	},
	Timeout: 3 * time.Second,
}

func TestContextCancellation(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// spawn 2 goroutines
	spawnRandomStatus(ctx, cancelFunc, &wg, ch)
	spawnDelay(ctx, cancelFunc, &wg, ch)

loop:
	for {
		select {
		case s := <-ch:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			fmt.Println("in main: canceled")
			break loop
		}
	}

	wg.Wait()
}

func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// return http.DefaultClient.Do(r)
	return client.Do(r)
}

// context.Context: interface
// sync.WaitGroup: struct
// chan string: reference type
func spawnRandomStatus(ctx context.Context, cancelFunc context.CancelFunc, wg *sync.WaitGroup, ch chan string) {
	go func() {
		defer wg.Done()

		for {
			// url := "https://httpbin.org/status/200,200,200,500"
			url := "http://localhost:81/status/200,200,200,500"
			resp, err := makeRequest(ctx, url)
			if err != nil {
				fmt.Println("error in status goroutine:", err)
				return
			}
			if resp.StatusCode == http.StatusInternalServerError {
				fmt.Println("bad status, exiting")
				cancelFunc()
				return
			}

			select {
			case ch <- "success from status":
			case <-ctx.Done():
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func spawnDelay(ctx context.Context, cancelFunc context.CancelFunc, wg *sync.WaitGroup, ch chan string) {
	go func() {
		defer wg.Done()

		for {
			// url := "https://httpbin.org/delay/1"
			url := "http://localhost:81/delay/1"
			resp, err := makeRequest(ctx, url)
			if err != nil {
				fmt.Println("error in delay goroutine:", err)
				cancelFunc()
				return
			}

			select {
			case ch <- "succcess from dealy: " + resp.Header.Get("date"):
			case <-ctx.Done():
			}
		}
	}()
}

// intterrupt long running tasks

// π/4 = 1 - 1/3 + 1/5 - 1/7 + 1/9 - ...
func leibnizMethod(ctx context.Context) (string, error) {
	var sum big.Float // 0
	sum.SetInt64(0)
	var d big.Float // 1
	d.SetInt64(1)

	two := big.NewFloat(2) // 2

	i := 0
	for {
		// check whether canceled
		if err := context.Cause(ctx); err != nil {
			result := sum.Text('g', 100)
			fmt.Println("cancelled after", i, "iterations: ", result)
			return result, err
		}

		// iteration
		var diff big.Float // 4
		diff.SetInt64(4)
		diff.Quo(&diff, &d) // 4 = 4 / d
		if i%2 == 0 {
			sum.Add(&sum, &diff) // sum = sum + diff
		} else {
			sum.Sub(&sum, &diff) // sum = sum - diff
		}
		d.Add(&d, two) // d = d + 2

		i++
	}
}

func TestCancelLongRunningTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pi, err := leibnizMethod(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("pi", pi)
		} else {
			fmt.Println("error:", err, "pi:", pi)
		}
	}

}
