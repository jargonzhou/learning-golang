package reflect_example

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckNilInterfaceValue(t *testing.T) {
	var a interface{}
	fmt.Println(a == nil, hasNoValue(a)) // true, true

	var b *int
	fmt.Println(b == nil, hasNoValue(b)) // true true

	var c interface{} = b                // pointer
	fmt.Println(c == nil, hasNoValue(c)) // false true

	var d int
	fmt.Println(hasNoValue(d)) // false

	var e interface{} = d
	fmt.Println(e == nil, hasNoValue(e)) // false false
}

func hasNoValue(i any) bool {
	iv := reflect.ValueOf(i)
	// IsValid reports whether v represents a value. It returns false if v is the zero Value.
	// If [Value.IsValid] returns false, all other methods except String panic.
	if !iv.IsValid() {
		return true
	}

	// IsNil reports whether its argument v is nil.
	// The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, IsNil panics.
	switch iv.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}
