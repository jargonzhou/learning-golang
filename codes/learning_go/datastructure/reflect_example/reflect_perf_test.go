package reflect_example

import (
	"learning_go/generics"
	"reflect"
	"testing"
)

var targetString []string
var targetInt []int

func Filter(slice any, filter any) any {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

// BenchmarkFilterString-20
//  1559976	       793.8 ns/op	     284 B/op	      14 allocs/op
// BenchmarkFilterInt-20
//  2166282	       573.7 ns/op	     195 B/op	      11 allocs/op
// BenchmarkGenericFilterString-20
// 16981676	        73.00 ns/op	     112 B/op	       3 allocs/op
// BenchmarkGenericFilterInt-20
// 36641892	        32.56 ns/op	      24 B/op	       2 allocs/op

func BenchmarkFilterString(b *testing.B) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetString = Filter(names, func(s string) bool {
			return len(s) > 3
		}).([]string)
	}
}

func BenchmarkFilterInt(b *testing.B) {
	ages := []int{20, 50, 13}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = Filter(ages, func(age int) bool {
			return age >= 18
		}).([]int)
	}
}

func BenchmarkGenericFilterString(b *testing.B) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetString = generics.Filter(names, func(s string) bool {
			return len(s) > 3
		})
	}
}

func BenchmarkGenericFilterInt(b *testing.B) {
	ages := []int{20, 50, 13}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = generics.Filter(ages, func(age int) bool {
			return age >= 18
		})
	}
}
