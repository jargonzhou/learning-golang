package types

import (
	"fmt"
	"testing"
)

// iota: assign an increasing value to a set of contants

type MailCategory int

const (
	UnCategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type BitField int

const (
	BitField1 BitField = 1 << iota
	BitField2
	BitField3
	BitField4
)

func TestIota(t *testing.T) {
	fmt.Println(UnCategorized, Personal, Spam, Social, Advertisements) // 0 1 2 3 4

	// the value of iota increments for each constant in the `const` block,
	// whether or not `iota` is used to define the value of a constant
	fmt.Println(Field1, Field2, Field3, Field4, Field5) // 0 2 20 20 4

	// %b: binary representation
	fmt.Printf("%b %b %b %b\n", BitField1, BitField2, BitField3, BitField4) // 1 10 100 1000
}
