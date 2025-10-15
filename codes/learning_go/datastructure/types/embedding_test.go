package types

import (
	"fmt"
	"testing"
)

//
// use embedding for composition
//

type SEmployee struct {
	Name string
	ID   string
}

func (e SEmployee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type SManager struct {
	SEmployee
	Reports []SEmployee
}

func (m SManager) FindNewEmployees() []SEmployee {
	return nil
}

type Inner struct {
	X int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2) // call methods on `Inner`
}

type Outer struct {
	Inner
	X int // same name in Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func TestComposition(t *testing.T) {
	m := SManager{
		SEmployee: SEmployee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []SEmployee{},
	}
	fmt.Println(m.ID)            // 12345
	fmt.Println(m.Description()) // Bob Bobson (12345)

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // 20
	fmt.Println(o.Inner.X) // 10
}

func TestEmbeddingIsNotInheritance(t *testing.T) {
	m := SManager{
		SEmployee: SEmployee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []SEmployee{},
	}
	// var eFail SEmployee = m // cannot use m (variable of struct type SManager) as SEmployee value in variable declaration
	var eOK SEmployee = m.SEmployee
	fmt.Println(eOK) // {Bob Bobson 12345}

	// Go has NO dynamic dispatch for concrete types
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		S: "Hello",
	}
	// call IntPrinter on Inner
	fmt.Println(o.Double()) // Inner: 20
}
