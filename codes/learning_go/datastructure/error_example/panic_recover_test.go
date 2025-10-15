package error_example

import (
	"fmt"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	// invalid operation: division by zero
	// fmt.Println(60 / 0)

	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

	// 60
	// 30
	// runtime error: integer divide by zero
	// 10
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
