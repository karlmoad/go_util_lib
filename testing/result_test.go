package testing

import (
	"github.com/karlmoad/go_util_lib/generics/result"
	"testing"
)

func TestNewResult(t *testing.T) {

	r := result.NewResult[int]()
	if r == nil {
		t.Errorf("NewResult[int]() returned nil")
	}

	if !r.Nothing() {
		t.Errorf("Result.Nothing returned false when true was expected")
	}
}

func TestResult_Valid(t *testing.T) {
	r := result.NewResult[int]()
	r.Set(20)

	if r.Nothing() {
		t.Errorf("Result.Nothing returned true when false was expected")
	}

	if r.Value() != 20 {
		t.Errorf("Result.Value() returned wrong value: got %v want %v", r.Value(), 20)
	}

	r2 := result.NewResultWithValue(100)

	if r2.Nothing() {
		t.Errorf("[#2]Result.Nothing returned true when false was expected")
	}

	if r2.Value() != 100 {
		t.Errorf("[#2]Result.Value() returned wrong value: got %v want %v", r.Value(), 100)
	}
}
