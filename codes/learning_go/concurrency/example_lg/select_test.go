package example_lg

import (
	"fmt"
	"testing"
)

// Example of turn off a case in select. - 'Learning Go': 12. Concurrency in Go, P.304
// read from or write to a nil channel: block.

func readFrom2Channels(in1, in2 chan int) []int {
	var out []int

	for count := 0; count < 2; { // 2 times: turn off cases
		select {
		case v, ok := <-in1:
			if !ok {
				fmt.Println(count, "turn off case <-in1")
				in1 = nil
				count++
				continue
			}
			// if success: no count++
			out = append(out, v)
		case v, ok := <-in2:
			if !ok {
				fmt.Println(count, "turn off case <-in2")
				in2 = nil
				count++
				continue
			}
			out = append(out, v)
		}
	}

	return out
}

func TestTurnoffSelectCase(t *testing.T) {
	in1 := make(chan int)
	in2 := make(chan int)
	go func() {
		for i := 10; i < 100; i += 10 { // 9 times
			in1 <- i
		}
		close(in1)
	}()

	go func() {
		for i := 20; i >= 0; i-- { // 21 times
			in2 <- i
		}
		close(in2)
	}()

	result := readFrom2Channels(in1, in2)
	fmt.Println(len(result), result)
}
