package mymath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Errorf("Add(3, 4) = %d; want 7", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 5)
	if result != 5 {
		t.Errorf("Subtract(10, 5) = %d; want 5", result)
	}
}
