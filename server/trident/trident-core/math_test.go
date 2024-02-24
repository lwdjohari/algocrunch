package tridentcore

import (
	"testing"
)

func TestAdd(t *testing.T) {

	sum := Foo(1, 2) // Using the Add function from core
	if sum != 3 {
		t.Errorf("Expected 3, got %d", sum)
	}
}
