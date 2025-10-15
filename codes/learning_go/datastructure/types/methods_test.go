package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// receiver: pointer, value
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func TestPerson(t *testing.T) {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	fmt.Println(p.String())
	// Fred Fredson, age 52
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func TestCounter(t *testing.T) {
	var c Counter
	fmt.Println(c.String())
	// total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	c.Increment() // (&c).Increment()
	fmt.Println(c.String())
	// total: 1, last updated: 2025-09-30 09:47:51.1196816 +0800 CST m=+0.003798301

	cp := &Counter{}
	fmt.Println(cp.String()) // (*cp).String()
	// total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	cp.Increment()
	fmt.Println(cp.String())
	// total: 1, last updated: 2025-09-30 09:47:51.1260466 +0800 CST m=+0.010163301
}

func TestCounterCornerCase(t *testing.T) {
	var c Counter
	doUpdateWrong(c)
	// in doUpdateWrong: total: 1, last updated: 2025-10-04 08:28:53.144606 +0800 CST m=+0.002382701
	fmt.Println("in test:", c.String())
	// in test: total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC

	doUpdateRight(&c)
	// in doUpdateRight: total: 1, last updated: 2025-10-04 08:28:53.149837 +0800 CST m=+0.007613701
	fmt.Println("in test:", c.String())
	// in test: total: 1, last updated: 2025-10-04 08:28:53.149837 +0800 CST m=+0.007613701

	assert := assert.New(t)
	assert.Panics(func() {
		((*Counter)(nil)).Increment()
	}, "panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]")
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

//
// methods for nil instances
//

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil { // nil receiver
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func TestMethodsForNilInstances(t *testing.T) {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}

//
// methods are functions:
// method value
// method expression
//

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func TestMethodsAreFunctions(t *testing.T) {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // 15

	// method value
	f1 := myAdder.AddTo
	fmt.Println(f1(10)) // 20

	// method expression
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // 25
}
