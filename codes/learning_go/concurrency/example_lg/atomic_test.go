package example_lg

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Example https://pkg.go.dev/sync/atomic#pkg-examples

// Config: program config updates and propagation

// mock load config: from file/db
func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func TestAtomicConfig(t *testing.T) {
	// setup config and initialize
	var config atomic.Value
	config.Store(loadConfig())

	// reload config periodly
	go func() {
		for {
			time.Sleep(2 * time.Second)
			config.Store(loadConfig())
		}
	}()

	// spawn worker goroutines
	num := 10
	var requestChannels []chan int
	for i := 0; i < num; i++ {
		rc := requests()
		requestChannels = append(requestChannels, rc)
		go func() {
			for r := range rc {
				c := config.Load()
				// handle request using latest config
				_, _ = r, c
				fmt.Println(i, r, c)
			}
		}()
	}

	for i := 0; i < num; i++ {
		for j := 0; j < 3; j++ {
			requestChannels[i] <- j
			if j == 1 {
				c := loadConfig()
				c["val"] = strconv.Itoa(j)
				config.Store(c)
			}
		}
		close(requestChannels[i])
	}
}

// ReadMostly

func TestReadMostly(t *testing.T) {
	var m atomic.Value
	m.Store(make(map[string]string))
	var mu sync.Mutex // for writers

	ctx, cancel := context.WithCancel(context.Background())

	read := func(key string) string {
		_m := m.Load().(map[string]string)
		return _m[key]
	}

	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()

		m1 := m.Load().(map[string]string)
		m2 := make(map[string]string)
		for k, v := range m1 { // copy
			m2[k] = v
		}
		m2[key] = val
		m.Store(m2)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				insert("hello", "golang "+strconv.Itoa(int(time.Now().UnixNano())))
				time.Sleep(30 * time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(read("hello"))
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
}
