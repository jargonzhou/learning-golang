package reflect_example

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Memoizer[T any](f T, expiration time.Duration) (T, error) {
	// function type
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		var zero T
		return zero, errors.New("only for functions")
	}

	argsType, err := makeArgsTypeStruct(ft)
	if err != nil {
		var zero T
		return zero, err
	}

	if ft.NumOut() == 0 {
		var zero T
		return zero, errors.New("must have at least one returned value")
	}

	// cache function result
	m := map[interface{}]outExp{}
	// function value
	fv := reflect.ValueOf(f)
	// create wapper function to memorize function result
	memo := reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		// argsType value
		iv := reflect.New(argsType).Elem()
		for k, v := range args {
			iv.Field(k).Set(v)
		}
		ivv := iv.Interface()

		ov, ok := m[ivv]
		now := time.Now()
		// if not cached, do calculation then cache result
		if !ok || ov.expiry.Before(now) {
			ov.out = fv.Call(args)
			ov.expiry = now.Add(expiration)
			m[ivv] = ov
		}
		return ov.out
	})
	return memo.Interface().(T), nil
}

func makeArgsTypeStruct(ft reflect.Type) (reflect.Type, error) { // ft is function
	if ft.NumIn() == 0 {
		return nil, errors.New("must have at lease on param")
	}

	// make struct fields
	sf := make([]reflect.StructField, 0, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		ct := ft.In(i)
		if !ct.Comparable() {
			return nil, fmt.Errorf("parameter %d of type %s and kind %v is not comparable", i+1, ct.Name(), ct.Kind())
		}

		sf = append(sf, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: ct,
		})
	}

	// make struct
	s := reflect.StructOf(sf)
	return s, nil
}

// output of function
type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

func AddSlowly(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}

func TestMemoizer(t *testing.T) {
	addSlowly, err := Memoizer(AddSlowly, 2*time.Second)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		start := time.Now()
		result := addSlowly(1, 2)
		end := time.Now()
		fmt.Println("got result", result, "in", end.Sub(start))
	}

	time.Sleep(3 * time.Second)
	start := time.Now()
	result := addSlowly(1, 2)
	end := time.Now()
	fmt.Println("got result", result, "in", end.Sub(start))
}
