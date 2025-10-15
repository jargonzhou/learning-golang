package pointer_example

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

// in Go runtime,
// a map is implemented as a pointer to a struct.
// a slice is implemented as a struct with 3 fileds:
//  length: int
//  capacity: int
//  a pointer to a block of memory

// pass a slice to function:
// modify slice's content reflected in original variable
// use `append` to change the length NOT reflect in original variable

func TestSliceAsFunctionArgument(t *testing.T) {
	s := []string{"1", "2", "3"}
	fmt.Println(s, len(s), cap(s), &(s[0])) // [1 2 3] 3 3 0xc000023500

	sliceArg1(s)
	fmt.Println(s, len(s), cap(s), &(s[0])) // [11 2 3] 3 3 0xc000023500

	sliceArg2(s)
	fmt.Println(s, len(s), cap(s), &(s[0])) // [11 2 3] 3 3 0xc000023500
}

func sliceArg1(s []string) {
	// modify content
	s[0] = "11"
}

func sliceArg2(s []string) {
	// use append()
	s = append(s, s...)
	fmt.Println("sliceArg2", s, len(s), cap(s), &(s[0])) // sliceArg2 [11 2 3 11 2 3] 6 6 0xc000080180
}

func TestSliceAsBuffer(t *testing.T) {
	processFileData("ptr_map_slice_test.go")
}

func processFileData(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data := make([]byte, 100) // slice buffer
	for {
		count, err := file.Read(data)
		fmt.Println(string(data[:count]))
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
}
