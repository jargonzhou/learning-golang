package example_lg

import "fmt"

// Example of a deadlock. - 'Learning Go': 12. Concurrency in Go, P.295
func DeadlockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromOuter := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromOuter)
	}()

	inOuter := 2
	ch2 <- inOuter
	fromGoroutine := <-ch1
	fmt.Println("main:", inOuter, fromGoroutine)
}

func DeadlockFixExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromOuter := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromOuter)
	}()

	inOuter := 2
	var fromGoroutine int
	// use select
	select {
	case ch2 <- inOuter:
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", inOuter, fromGoroutine)
}
