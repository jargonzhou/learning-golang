package example_cig

import (
	"fmt"
	"sync"
)

// Example of closure capture. - 'Concurrency in Go' P.41
// Things changed in loop closure capture.
//
// Output
//
//	good day
//	greetings
//	hello
func ClosureExample() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
