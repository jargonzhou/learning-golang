package tools

import (
	"fmt"
	"os"
	"os/signal"
)

// Create new singal handler.
// sig example: syscall.SIGINT, syscall.SIGTERM
func NewSignalHandler(done chan<- bool, signals ...os.Signal) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)

	go func() {
		// block until a signal is received
		sig := <-sigs
		fmt.Println("Receive singal: ", sig)

		// do some cleanup
		fmt.Println("Tear down")

		done <- true
	}()
}
