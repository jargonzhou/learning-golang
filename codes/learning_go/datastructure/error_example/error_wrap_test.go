package error_example

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// wrap: Errorf() with %w
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func TestErrorWrapping(t *testing.T) {
	err := fileChecker("not-found.txt")
	if err != nil {
		fmt.Println(err)
		// in fileChecker: open not-found.txt: The system cannot find the file specified.
		// unwrap
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
			// open not-found.txt: The system cannot find the file specified.
		}
		// Is reports whether any error in err's tree matches target. use == to compare
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file doesn't exist")
		}
	}
}

type WrappedStatusErr struct {
	Status  Status
	Message string
	Err     error // inner error
}

func (wse WrappedStatusErr) Error() string {
	return wse.Message
}

func (wse WrappedStatusErr) Unwrap() error {
	return wse.Err
}

func LoginAndGetData2(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, WrappedStatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}

	data, err := getData(token, file)
	if err != nil {
		return nil, WrappedStatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

//
// wrap multiple errors:
// errors.Join()
// fmt.Errorf(), multiple %w verbs
// `Errors []error` in customed error struct

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p Person) error {
	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}
	// use errors.Join()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func TestWrapMultipleErrors(t *testing.T) {
	assert := assert.New(t)

	p := Person{
		FirstName: "first",
		LastName:  "",
		Age:       -1,
	}
	err := ValidatePerson(p)
	fmt.Println(err)
	// field LastName cannot be empty
	// field Age cannot be negative
	assert.Nil(errors.Unwrap(err))

	err1 := errors.New("first error")
	err2 := errors.New("second error")
	err3 := errors.New("third error")
	err = fmt.Errorf("first: %w, second: %w, third: %w", err1, err2, err3)
	fmt.Println(err)
	// first: first error, second: second error, third: third error
	assert.Nil(errors.Unwrap(err))
}

type MyError struct {
	Code   int
	Codes  []int
	Errors []error
}

func (m MyError) Error() string {
	return errors.Join(m.Errors...).Error()
}

func (m MyError) Unwrap() []error {
	return m.Errors
}

func (m MyError) Is(target error) bool {
	// override default == in Is
	if me, ok := target.(MyError); ok {
		return slices.Equal(me.Codes, m.Codes)
	}
	return false
}

func TestHandleMultipleErrors(t *testing.T) {
	var err error
	p := Person{
		FirstName: "first",
		LastName:  "",
		Age:       -1,
	}
	err = ValidatePerson(p)
	switch err := err.(type) {
	case interface{ Unwrap() error }: // anonymous interface
		innerErr := err.Unwrap()
		fmt.Println("ERROR:", innerErr)
	case interface{ Unwrap() []error }:
		innerErrs := err.Unwrap()
		for i, innerErr := range innerErrs {
			fmt.Println("ERROR", i, ":", innerErr)
		}
	default:
		fmt.Println(err)
	}
	// ERROR 0 : field LastName cannot be empty
	// ERROR 1 : field Age cannot be negative
}

//
// Is(), As()
//

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) CodeVals() int {
	return re.Code
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code

		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func TestIs(t *testing.T) {
	err := ResourceErr{
		Resource: "database",
	}
	target := fmt.Errorf("Resource error: %w", err)
	if errors.Is(target, ResourceErr{Resource: "database"}) {
		fmt.Println("The database is broken:", target)
		// The database is broken: Resource error: database: 0
	}
}

func TestAs(t *testing.T) {
	err := ResourceErr{
		Resource: "database",
		Code:     42,
	}
	target := fmt.Errorf("Resource error: %w", err)
	var re ResourceErr
	// As finds the first error in err's tree that matches target, and if one is found,
	// sets target to that error value and returns true. Otherwise, it returns false.
	if errors.As(target, &re) {
		fmt.Println(re.Code) // 42
	}

	// error match a interface
	var coder interface {
		CodeVals() int
	}
	if errors.As(target, &coder) {
		fmt.Println(coder.CodeVals()) // 42
	}
}

func TestWrappingErrorWithDefer(t *testing.T) {
	v, err := doSomething(3, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	// in doSomething: t=2 error
}

func doSomething(val1, val2 int) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in doSomething: %w", err)
		}
	}()

	val3, err := doThing(1, val1)
	if err != nil {
		return 0, err
	}
	val4, err := doThing(2, val2)
	if err != nil {
		return 0, err
	}
	return doThing(val3, val4)
}

func doThing(t int, val int) (int, error) {
	if t == 2 {
		return 0, errors.New("t=2 error")
	}
	return val, nil
}
