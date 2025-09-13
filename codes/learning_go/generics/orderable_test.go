package generics

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Orderable interface {
	// any
	Order(any) int
}

type Tree struct {
	val         Orderable
	left, right *Tree
}

func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil {
		return &Tree{val: val}
	}

	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}

	return t
}

type OrderableInt int

func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
}

type OrderableString string

func (os OrderableString) Order(val any) int {
	return strings.Compare(string(os), val.(string))
}

func TestOrderable(t *testing.T) {
	var it *Tree
	it = it.Insert(OrderableInt(5))
	fmt.Printf("%v\n", it)

	assert := assert.New(t)
	errString := "panic: interface conversion: interface {} is generics.OrderableInt, not string [recovered, repanicked]"
	assert.Panics(func() {
		it = it.Insert(OrderableString("nope"))
		fmt.Printf("%v\n", it)
	}, errString)
	// assert.PanicsWithError(errString, func() {
	// 	it = it.Insert(OrderableString("nope"))
	// 	fmt.Printf("%v\n", it)
	// })
}
