package calculator_test

import (
	"testing"
	"ztg/calculator"
)

func assertEqual(t *testing.T, want, got any) {
	t.Helper()

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestAddNumbers(t *testing.T) {
	result := calculator.Add(1, 2)

	assertEqual(t, 3, result)
}

func TestAddOnlyPositive(t *testing.T) {
	result := calculator.Add(1, -2)

	assertEqual(t, 0, result)
}
