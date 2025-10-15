package example_cig

import "fmt"

// Example of data race.
func RaceExample() {
	var data int
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
