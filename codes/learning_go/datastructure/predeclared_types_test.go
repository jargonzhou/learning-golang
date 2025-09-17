// Examples of 'Learning Go': 2. Predeclared Types and Declarations
//
// zero value
// literals: integer, floating-point, rune, string(intepreted string, raw string), complex numbers
// booleans: bool
// integers: int8, int16, int32, int64, uint8, uint16, uint32, uint64, byte, int, uint, rune, uintptr
// floats: float32, float64, Inf, NaN
// complex numbers: complex64, complex128
// strings
package datastructure

import (
	"fmt"
	"math/cmplx"
	"testing"
)

func TestComplexNumbers(t *testing.T) {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

func TestRune(t *testing.T) {
	var myFirstInitial rune = 'J' // good - the type name matches the usage
	var myLastInitial int32 = 'B' // bad - legal but confusing
	fmt.Println(myFirstInitial)
	fmt.Println((myLastInitial))
}

func TestExplicitTypeConversion(t *testing.T) {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	// integer
	// var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, sum4)
}

func TestLiterals(t *testing.T) {
	// literals are untyped
	var x float64 = 10
	var y float64 = 200.3 * 5
	fmt.Println(x)
	fmt.Println(y)
}
