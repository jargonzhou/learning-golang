package tests_test

import (
	"testing"

	"learning_go/tests"
)

func TestAddNumber(t *testing.T) {
	result := tests.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expect 5 but got", result)
	}
}
