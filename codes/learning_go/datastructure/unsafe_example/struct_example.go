package unsafe_example

import "fmt"

type HasUnexportedField struct {
	A int
	b bool // unexported field
	C string
}

func (huf *HasUnexportedField) ToString() string {
	return fmt.Sprintf("HasUnexportedField{A=%d,b=%t,C=%s}\n", huf.A, huf.b, huf.C)
}
