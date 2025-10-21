package likes

import (
	"testing"
)

func clear() {
	for k := range m {
		delete(m, k)
	}
}

func check(t *testing.T, key string, ids []int) {
	t.Helper()

	v, ok := m[key]
	if !ok {
		t.Fatal("key was not found in map")
	} else if len(v) != len(ids) {
		t.Fatalf("expected len of %d, got %d", len(ids), len(v))
	}

	for i := range ids {
		if v[i] != ids[i] {
			t.Errorf("expected index %d to be %d, got %d", i, ids[i], v[i])
		}
	}
}
func TestAddLikes(t *testing.T) {
	clear()

	add("test", 123)
	add("test", 456)

	check(t, "test", []int{123, 456})
}

func TestAddDuplicate(t *testing.T) {
	clear()

	add("test", 123)
	add("test", 123)

	check(t, "test", []int{123})
}

func TestUnlike(t *testing.T) {
	clear()

	add("test", 123)
	add("test", 456)
	add("test", 789)
	add("test", 321)
	remove("test", 456)

	check(t, "test", []int{123, 789, 321})

}
