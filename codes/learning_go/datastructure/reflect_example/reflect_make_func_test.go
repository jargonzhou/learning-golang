package reflect_example

import (
	"fmt"
	"testing"
	"time"
)

func TestMakeTimedFunction(t *testing.T) {
	timed := MakeTimedFunction(timeMe).(func(int) int)
	fmt.Println(timed(2))
}

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	return a * 2
}
