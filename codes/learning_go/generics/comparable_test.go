// All your comparable types: https://go.dev/blog/comparable
package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("equal!")
	}
}

func TestComparable(t *testing.T) {
	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	var a4 Thinger = a2 // assign ThingerInt to Thinger
	var b4 Thinger = b2
	Comparer(a4, b4)

	assert := assert.New(t)
	assert.Panics(func() {
		var a3 ThingerSlice = []int{1, 2, 3}
		var b3 ThingerSlice = []int{1, 2, 3}
		var a4 Thinger = a3 // assign ThingerSlice to Thinger
		var b4 Thinger = b3
		Comparer(a4, b4)
	})

	// var a3 ThingerSlice = []int{1, 2, 3}
	// var b3 ThingerSlice = []int{1, 2, 3}
	// // ThingerSlice does not satisfy comparable
	// Comparer(a3, b3)
}
