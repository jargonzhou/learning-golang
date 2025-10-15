package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Incrementer interface {
	Increment()
}

func TestInterfaces(t *testing.T) {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer

	// Counter: methods_test.go
	pCounter := &Counter{}
	vCounter := Counter{}

	myStringer = pCounter
	myStringer = vCounter

	myIncrementer = pCounter
	// ERROR: cannot use vCounter (variable of struct type Counter) as Incrementer value in assignment: Counter does not implement Incrementer (method Increment has pointer receiver)
	// myIncrementer = vCounter

	_ = myStringer
	_ = myIncrementer
}

//
// interface are type-safe duck typing
//

type ILogic interface {
	Process(data string) string
}

type SLogicProvider struct{}

func (lp SLogicProvider) Process(data string) string {
	return "LP " + data
}

type SClient struct {
	L ILogic
}

func (c SClient) Program(data string) {
	fmt.Println(c.L.Process(data))
}

func TestTypeSafeDuckTyping(t *testing.T) {
	c := SClient{
		L: SLogicProvider{},
	}
	c.Program("Hallo")
	// LP Hallo
}

//
// accept interfaces, return structs
//

//
// nil
// interface runtime representation: a struct with 2 pointer fields:
//  one for value
//  one for the type of the value

func TestNil(t *testing.T) {
	var pCounter *Counter
	var incrementer Incrementer

	fmt.Println(pCounter == nil)    // true
	fmt.Println(incrementer == nil) // true

	incrementer = pCounter
	fmt.Println(incrementer == nil) // false
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

// slices are NOT comparable
type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func TestInterfacesAreComparable(t *testing.T) {
	assert := assert.New(t)

	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}

	DoublerCompare(&di, &di2) // false
	DoublerCompare(&di, dis)  // false
	assert.Panics(func() {
		DoublerCompare(dis, dis2)
	}, "panic: runtime error: comparing uncomparable type types.DoubleIntSlice [recovered, repanicked]")
}

func TestEmptyInterface(t *testing.T) {
	// the empty interface: interface{}
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	_ = i

	// any is type alias for interface{}
	var i2 any
	fmt.Println(i2 == nil) // true
	_ = i
}
