package finder

import (
	"errors"
	"testing"
)

func TestFindTheItem(t *testing.T) {
	values := []string{"a", "b", "c"}
	idx, err := Find(values, "b")
	if err != nil {
		t.Fatal(err)
	} else if idx != 1 {
		t.Errorf("expected index to be 2, got %d", idx)
	}
}

func TestFindItemDifferentCasing(t *testing.T) {
	values := []string{"a", "b", "c"}
	idx, err := Find(values, "B")
	if err != nil {
		t.Fatal(err)
	} else if idx != 1 {
		t.Errorf("expected index to be 2, got %d", idx)
	}
}

func TestFindMissingItemError(t *testing.T) {
	values := []string{"a", "b", "c"}
	if _, err := Find(values, "z"); err != nil {
		if !errors.Is(err, ErrItemNotFound) {
			t.Errorf("expected error typed ErrItemNotFound")
		}
		return
	}

	t.Error("expected an error returned")
}

func TestFindWithEmptySliceReturnError(t *testing.T) {
	if _, err := Find(nil, "a"); err != nil {
		if !errors.Is(err, ErrInvalidSlice) {
			t.Error("expected error to be ErrInvalidSlice")
		}
		return
	}

	t.Errorf("expected to return an error for nil slice")
}
