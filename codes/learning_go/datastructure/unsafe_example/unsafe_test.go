package unsafe_example

import (
	"fmt"
	"testing"
	"unsafe"
)

type BoolInt struct {
	b bool
	i int64
}

type IntBool struct {
	i int64
	b bool
}

type BoolIntBool struct {
	b  bool
	i  int64
	b2 bool
}
type BoolBoolInt struct {
	b  bool
	b2 bool
	i  int64
}
type IntBoolBool struct {
	i  int64
	b  bool
	b2 bool
}

func TestSizeofAndOffsetof(t *testing.T) {
	// bool, int64

	bi := BoolInt{}
	fmt.Println(unsafe.Sizeof(bi), unsafe.Offsetof(bi.b), unsafe.Offsetof(bi.i)) // 16 0 8

	ib := IntBool{}
	fmt.Println(unsafe.Sizeof(ib), unsafe.Offsetof(ib.i), unsafe.Offsetof(ib.b)) // 16 0 8

	// 2 bools, int64
	bib := BoolIntBool{}
	bbi := BoolBoolInt{}
	ibb := IntBoolBool{}
	// 24 0 8 16
	fmt.Println(unsafe.Sizeof(bib), unsafe.Offsetof(bib.b), unsafe.Offsetof(bib.i), unsafe.Offsetof(bib.b2))
	// 16 0 1 8
	fmt.Println(unsafe.Sizeof(bbi), unsafe.Offsetof(bbi.b), unsafe.Offsetof(bbi.b2), unsafe.Offsetof(bbi.i))
	// 16 0 8 9
	fmt.Println(unsafe.Sizeof(ibb), unsafe.Offsetof(ibb.i), unsafe.Offsetof(ibb.b), unsafe.Offsetof(ibb.b2))
}
