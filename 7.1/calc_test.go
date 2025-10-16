package calculator

import (
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("wanted %d, got %d", 3, result)
	}
}
