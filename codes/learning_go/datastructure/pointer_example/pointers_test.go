package pointer_example

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	var x int32 = 10
	var y bool = true

	pointerX := &x
	pointerY := &y
	var pointerZ *string

	fmt.Printf("x: value=%d address=%p\n", *pointerX, pointerX) // x: value=10 address=0xc00000aad8
	fmt.Printf("y: value=%t address=%p\n", *pointerY, pointerY) // y: value=true address=0xc00000aadc
	if pointerZ != nil {
		fmt.Printf("z: value=%s address=%p\n", *pointerZ, pointerZ)
	} else {
		fmt.Printf("z: address=%p\n", pointerZ) // z: address=0x0
	}
	fmt.Printf("nil pointer address=%p\n", (*string)(nil)) // nil pointer address=0x0

	assert := assert.New(t)
	assert.Nil(pointerZ)
	assert.Panics(func() {
		fmt.Println(*pointerZ)
	}, "panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]")

	// new()
	var xx = new(int)
	assert.NotNil(xx)
	assert.Equal(0, *xx)
}

type person struct {
	FirstName  string
	MiddleName *string // pointer type field
	LastName   string
}

func MakePointer[T any](t T) *T {
	return &t
}

func TestPerson(t *testing.T) {
	p := person{
		FirstName:  "Pat",
		MiddleName: MakePointer("Perry"),
		LastName:   "Peterson",
	}
	fmt.Printf("%#v\n%v\n", p, p)
}

func TestPassingPointerArgument(t *testing.T) {
	assert := assert.New(t)

	var f *int
	assert.Nil(f)

	failedUpdate(f)
	assert.Nil(f) // still nil

	// f = new(int)
	// *f = 10
	x := 10
	f = &x
	update(f)
	assert.Equal(20, *f)
}

func failedUpdate(g *int) { // pass a copy of pointer
	x := 10
	g = &x
}

func update(g *int) {
	*g = 20
}

//
// style
//

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

func MakeFooPreferred() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func TestPointerArgSty(t *testing.T) {
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f) // 2nd paramenter expect pointer
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", f)
	// struct { Name string "json:\"name\""; Age int "json:\"age\"" }{Name:"Bob", Age:30}
}
