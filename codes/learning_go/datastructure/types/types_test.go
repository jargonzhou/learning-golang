package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// struct type
type ExamplePerson struct {
	FirstName string
	LastName  string
	Age       int
}

// alias: with underlying type
type Score int
type Converter func(string) Score
type TeamScores map[string]Score

//
// type declaration are NOT inheritance
//

type HighScore Score
type Employee ExamplePerson

func TestTyp(t *testing.T) {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	// SHOULD use type conversion
	// hs = s // ERROR: cannot use s (variable of int type Score) as HighScore value in assignment
	// s = i // ERROR: cannot use i (variable of type int) as Score value in assignment
	s = Score(i)
	hs = HighScore(s)

	// underlying types are built-in types
	scoreWithBonus := s + 100
	fmt.Println(hs, scoreWithBonus)
}

//
// type assertions
// type switches
//

type MyInt int

func TestTypeAssertion(t *testing.T) {
	assert := assert.New(t)

	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)     // type assertions
	fmt.Println(i2 + 1) // 21

	// panics
	assert.Panics(func() {
		i3 := i.(string)
		_ = i3
	}, "panic: interface conversion: interface {} is types.MyInt, not string [recovered, repanicked]")
	assert.Panics(func() {
		i3 := i.(int)
		_ = i3
	}, "panic: interface conversion: interface {} is types.MyInt, not int [recovered, repanicked]")

	i3, ok := i.(int)
	if !ok {
		fmt.Printf("unexpected type for %v\n", i) // unexpected type for 20
	}
	_ = i3
}

func TestTypeSwitch(t *testing.T) {
	var i any
	i = nil
	doThings(i) // <nil> nil: (<nil>)

	i = 42
	doThings(i) // int int: (int)

	i = MyInt(10)
	doThings(i) // types.MyInt MyInt: (types.MyInt)

	i = false
	doThings(i) // bool bool, rune: (bool)

	i = "42"
	doThings(i) // string ???: (string)
}

func doThings(i any) {
	fmt.Printf("%T ", i)
	// .(type)
	switch t := i.(type) {
	case nil:
		fmt.Printf("nil: (%T)\n", t)
	case int:
		fmt.Printf("int: (%T)\n", t)
	case MyInt:
		fmt.Printf("MyInt: (%T)\n", t)
	case bool, rune:
		fmt.Printf("bool, rune: (%T)\n", t)
	default:
		fmt.Printf("???: (%T)\n", t)
		_ = t
	}
}
