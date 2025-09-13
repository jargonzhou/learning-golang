package tests

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expect 5 but got", result)
	}
}
