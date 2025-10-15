package unsafe_example_test

import (
	"fmt"
	"learning_go/datastructure/unsafe_example"
	"reflect"
	"testing"
	"unsafe"
)

func SetBUnsafe(huf *unsafe_example.HasUnexportedField) {
	// struct field
	sf, _ := reflect.TypeOf(huf).Elem().FieldByName("b")
	// read old value
	offset := sf.Offset
	start := unsafe.Pointer(huf)
	pos := unsafe.Add(start, offset)
	b := (*bool)(pos)
	fmt.Println(*b)
	// write new value
	*b = true
}

func TestSetBUnsafe(t *testing.T) {
	huf := &unsafe_example.HasUnexportedField{
		A: 42,
		C: "42",
	}
	fmt.Println(huf.ToString()) // HasUnexportedField{A=42,b=false,C=42}

	SetBUnsafe(huf)
	fmt.Println(huf.ToString()) // HasUnexportedField{A=42,b=true,C=42}
}
