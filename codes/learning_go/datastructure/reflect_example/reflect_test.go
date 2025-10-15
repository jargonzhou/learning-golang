package reflect_example

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

var (
	stringType      = reflect.TypeOf((*string)(nil)).Elem()
	stringSliceType = reflect.TypeOf([]string(nil))

	intType = reflect.TypeOf(0)
)

func TestTypeKindValue(t *testing.T) {
	// reflect.Type
	// Name(), String(),
	// Kind()
	// Elem(), NumField(), Field()

	TestInt(t)

	TestStruct(t)

	TestStringSlice(t)

	TestFunc(t)
}

func TestInt(t *testing.T) {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name()) // int

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.String()) // *int
	// reflect.Kind
	xptk := xpt.Kind()
	fmt.Println(xptk.String())                        // ptr
	fmt.Println(xpt.Elem().Name(), xpt.Elem().Kind()) // int int

	// reflect.Value
	xv := reflect.ValueOf(x)
	fmt.Println(xv) // 0
	if xv.CanInt() {
		fmt.Println(xv.Int()) // 0
	}
}

func TestStruct(t *testing.T) {
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name()) // Foo
	for i := 0; i < ft.NumField(); i++ {
		// reflect.StructField
		field := ft.Field(i)
		fmt.Println(field.Name, field.Type.Name(), field.Tag.Get("myTag"))
	}
	// A int value
	// B string value2
}

func TestStringSlice(t *testing.T) {
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	fmt.Println(sv) // [a b c]
	s2 := sv.Interface().([]string)
	fmt.Println(s2) // [a b c]
}

func TestFunc(t *testing.T) {
	f := func(in int) string {
		return strconv.Itoa(in)
	}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.String()) // func(int) string

	// invoke function
	fv := reflect.ValueOf(f)
	var in []reflect.Value
	in1 := reflect.New(intType).Elem()
	in1.SetInt(1)
	in = append(in, in1)
	res := fv.Call(in)
	fmt.Println(len(res), res) // 1 [1]
	result := res[0]
	fmt.Println(result.String()) // 1
}

func TestSetVariableValue(t *testing.T) {
	i := 10
	// 1. pass a pointer
	iv := reflect.ValueOf(&i)
	// 2. get the actual value to set
	ivv := iv.Elem()
	// 3. set the value
	ivv.SetInt(20)
	fmt.Println(i) // 20
}

// Kind
//
// Invalid Kind = iota
// Bool
// Int
// Int8
// Int16
// Int32
// Int64
// Uint
// Uint8
// Uint16
// Uint32
// Uint64
// Uintptr
// Float32
// Float64
// Complex64
// Complex128
// Array
// Chan
// Func
// Interface
// Map
// Pointer
// Slice
// String
// Struct
// UnsafePointer

func TestMakeNewValues(t *testing.T) {
	// reflect.Type String()
	fmt.Println(stringType.String()) // string

	fmt.Println(stringSliceType.String()) // []string

	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")
	ssv = reflect.Append(ssv, sv)
	// reflect.Value String()
	fmt.Println(ssv.String(), ssv) // <[]string Value> [hello]

	ss := ssv.Interface().([]string)
	fmt.Println(ss) // [hello]
}
