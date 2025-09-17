// Examples of 'Learning Go': 3. Composite Types
//
// arrays
// slices
// strings, runes, bytes: package strings, unicode/utf8
// maps
// structs
package datastructure

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	//	var x [3]int
	// var x = [3]int{10, 20, 30}
	// var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	// var x = [...]int{10, 20, 30}
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
	assert := assert.New(t)
	assert.Equal(x, y)

	// multidimensional arrays
	// var x [2][3]int

	// element access
	x[0] = 10
	fmt.Println(x[2])

	// length
	fmt.Println(len(x))
}

func TestSclice(t *testing.T) {
	// var x = []int{10, 20, 30}
	var x = []int{1, 5: 4, 6, 10: 100, 15}
	// var x [][]int

	x[0] = 10
	fmt.Println(x[2])

	// not comparable
	var s []int
	assert := assert.New(t)
	assert.True(s == nil)
}

func TestSlicesPackage(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	// ERROR: in call to slices.Equal, type []string of s does not match inferred type []int for S
	// fmt.Println(slices.Equal(x, s))
	fmt.Println(s)
}

func TestBIFForSlices(t *testing.T) {
	// len
	x := []int{1, 2, 3, 4, 5}
	assert := assert.New(t)
	assert.Equal(5, len(x))

	// append
	x = append(x, 10)
	x = append(x, 5, 6, 7)
	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x)

	// cap
	x = nil
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// make
	x = make([]int, 5)
	x = append(x, 10)
	fmt.Println(x)

	// clear
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))
}

func TestSliceDeclaration(t *testing.T) {
	assert := assert.New(t)

	// nil
	var data []int
	assert.True(data == nil)

	// empty slice literal
	var x = []int{}
	assert.False(x == nil)

	// with default values
	data2 := []int{2, 4, 6, 8}

	fmt.Println(data, x, data2)

	// copy
	x = []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num) // [1 2 3 4] 4
}

func TestSlicingSlice(t *testing.T) {
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// d := x[1:3]
	// e := x[:]
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// fmt.Println("d:", d)
	// fmt.Println("e:", e)

	// slices share memory
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// x[1] = "y"
	// y[0] = "x"
	// z[1] = "z"
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// odd behaviour with append
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y)) // 4 4
	y = append(y, "z")
	fmt.Println(cap(x), cap(y)) // 4 4
	fmt.Println("x:", x)        // x: [a b z d]
	fmt.Println("y:", y)        // y: [a b z]
}

func TestConversionBetweenArrayAndSlice(t *testing.T) {
	// array to slice
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// slice to array: use type converion
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)

	assert := assert.New(t)
	assert.Panics(func() {
		panicArray := [5]int(xSlice)
		fmt.Println(panicArray)
	})
}

func TestStringRuneBytes(t *testing.T) {
	var s string = "Hello 🌞"
	var b byte = s[6]
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(b, s, s2, s3, s4)

	// return length in bytes
	fmt.Println(len(s)) // 10
}

func TestRuneByteToString(t *testing.T) {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)

	fmt.Println(s, s2)
}

func TestStringToSlice(t *testing.T) {
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
