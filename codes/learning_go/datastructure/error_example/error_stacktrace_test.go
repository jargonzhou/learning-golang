package error_example

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/errors" // cockroachdb/errors
)

func returnError() (string, error) {
	return "", errors.New("something wrong")
}

func TestErrorStackTrace(t *testing.T) {
	v, err := returnError()
	if err != nil {
		// use %+v
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(v)

	//	something wrong
	// (1) attached stack trace
	//	-- stack trace:
	//	| learning_go/datastructure/error_example.returnError
	//	| 	.../learning_go/datastructure/error_example/error_stacktrace_test.go:11
	//	| learning_go/datastructure/error_example.TestErrorStackTrace
	//	| 	.../learning_go/datastructure/error_example/error_stacktrace_test.go:15
	//	| testing.tRunner
	//	| 	D:/software/go1.25.1/src/testing/testing.go:1934
	//	| runtime.goexit
	//	| 	D:/software/go1.25.1/src/runtime/asm_amd64.s:1693
	// Wraps: (2) something wrong
	// Error types: (1) *withstack.withStack (2) *errutil.leafError
}
