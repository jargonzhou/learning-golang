package generics

// type term

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// can only used as type constraints
type Integer interface {
	// type element: compose of type terms
	// exact match
	// int | int8 | int16 | int32 | int64 |
	// uint | uint8 | uint16 | uint32 | uint64 | uintptr

	// consider underlying type
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func divAndReainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	// operators: /, %
	return num / denom, num % denom, nil
}

type MyInt int

// both type elements and method elements in interface used for type parameter
type PrintableInt interface {
	~int
	String() string
}

// declare a type parameter interface that cannot instantiate
type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

// type terms can also be:
// slices, maps, arrays, channels, structs, functions
func TestTypeTerm(t *testing.T) {
	var a uint = 18_446_744_073_709_551_615
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(divAndReainder(a, b))

	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(divAndReainder(myA, myB))

	// ERROR: int does not satisfy ImpossiblePrintableInt (missing method String)
	// s := ImpossibleStruct[int]{10}
	// ERROR: MyInt does not satisfy ImpossiblePrintableInt (possibly missing ~ for int in ImpossiblePrintableInt)
	// s2 := ImpossibleStruct[MyInt]{10}
	// fmt.Printf("%v %v\n", s, s2)
}

// type inference
type Integer2 interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer2](in T1) T2 {
	return T2(in)
}

func TestTypeInference(t *testing.T) {
	var a int = 10
	// cannot infer return type: must specify all type argument
	b := Convert[int, int64](a)
	fmt.Println(b)
}

// limitation

// cannot convert 1_000 (untyped int constant 1000) to type T
// func Plus1000[T Integer2](in T) T {
// 	return in + 1_000
// }

func Plus100[T Integer2](in T) T {
	return in + 100
}
