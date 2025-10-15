package main

import (
	"fmt"
	"learning_go/concurrency/example_cig"
	"learning_go/concurrency/example_lg"
	"learning_go/tools"
	"syscall"
	"time"
)

// fatal error: all goroutines are asleep - deadlock!
func DeadlockExample() {
	// concurrency.DeadlockExample()

	example_lg.DeadlockFixExample()
}

func PanicExample() {
	go func() {
		panic("test panic")
	}()
}

// Wait sigal for iterrupt.
func SignalExample() {
	done := make(chan bool, 1)

	tools.NewSignalHandler(done, syscall.SIGINT, syscall.SIGTERM)
	tools.StartPprof("localhost:6060")

	<-done
	fmt.Println("Done!")
}

// Run with: go run --race main.go
func RaceDetectorExample() {
	example_cig.RaceExample()
}

func PProfExample() {
	tools.SetUpLog()
	tools.DumpProfiles(tools.ProfileGoroutine, 30*time.Second)
}

func ClosureExample() {
	example_cig.ClosureExample()
}

func main() {
	fmt.Println("Hello, 世界!")

	DeadlockExample()

	// PanicExample()
	// RaceDetectorExample()
	// PProfExample()
	// ClosureExample()

	SignalExample()
}
