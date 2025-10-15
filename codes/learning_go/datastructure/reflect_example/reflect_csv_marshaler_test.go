package reflect_example

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCSVMarshaler(t *testing.T) {
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"
`

	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	Unmarshal(allData, &entries)
	for _, entry := range entries {
		fmt.Printf("%#v\n", entry)
	}

	fmt.Println()

	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

// Marshal maps all of the structs in a slice of structs to a slice of slice
// of strings.
// The first row written is the header with the column names.
func Marshal(v any) ([][]string, error) { // v is slice of structs
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}

	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(vt reflect.Type) []string { // vt is struct
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) { // vv is struct
	var row []string
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}

		// process fields in struct
		switch fieldVal.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}

	return row, nil
}

// Unmarshal maps all of the rows of data in a slice of slice of strings
// into a slice of structs.
// The first row is assumed to be the header with the column names.
func Unmarshal(data [][]string, v any) error { // v is pointer to slice of structs
	sliceValPointer := reflect.ValueOf(v)
	if sliceValPointer.Kind() != reflect.Pointer {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPointer.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of structs")
	}

	// header
	header := data[0]
	namePos := make(map[string]int, len(header))
	for i, name := range header {
		namePos[name] = i
	}
	// rows
	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem() // make new struct
		err := unmarshalOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error { // vv is struct
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeFiled := vt.Field(i)
		pos, ok := namePos[typeFiled.Tag.Get("csv")]
		if !ok {
			continue
		}

		val := row[pos]
		field := vv.Field(i)
		switch field.Kind() {
		case reflect.Int:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}

func MakeTimedFunction(f any) any { // f is function
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		out := fv.Call(args)
		end := time.Now()
		fmt.Println("last", end.Sub(start))
		return out
	})
	// Interface returns v's current value as an interface{}.
	return wrapperF.Interface()
}
