// Examples of 'Learning Go': 2. Predeclared Types and Declarations
//
// var
// :=: short declaration and assignment
// const: typed, untyped
package datastructure

import (
	"fmt"
	"testing"
)

// declaration list
var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "hello"
	f, g string
)

// constants
// typed
const cx int64 = 10

// untyped
const (
	cidKey   = "id"
	cnameKey = "name"
)
const cz = 20 * 10

func TestDeclarationList(t *testing.T) {
	fmt.Println(x, y, z, d, e, f, g)
}

func TestConst(t *testing.T) {
	const cy = "hello"
	fmt.Println(cx)
	fmt.Println(cy)
	fmt.Println(cidKey, cnameKey, cz)

	// ERROR: cannot assign to cx (neither addressable nor a map index expression)
	// cx = cx + 1
	// ERROR: cannot assign to cy (neither addressable nor a map index expression)
	// cy = "bye"
}

func TestVariableNames(t *testing.T) {
	// Variable names you should never use
	_0 := 0_0
	_1 := 20
	π := 3
	ａ := "hello"              // Unicode U+FF41
	__ := "double underscore" // two underscores
	fmt.Println(_0)
	fmt.Println(_1)
	fmt.Println(π)
	fmt.Println(ａ)
	fmt.Println(__)
}
