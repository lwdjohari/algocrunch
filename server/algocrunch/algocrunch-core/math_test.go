package algocrunchcore

import (
	"testing"
)

func TestAdd(t *testing.T) {

	sum := Foo(1, 2) // Using the Add function from core
	if sum.Unwrap() != 3 {
		t.Errorf("Expected 3, got %d", sum)
	}
}
