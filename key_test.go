package irmin

import (
	"testing"
)

func TestKeyRoundtrip(t *testing.T) {
	key0 := NewKey("/a/b/c")
	key1 := NewKey("a/b/c/")
	key2 := Key{"a", "b", "c"}

	a := key0.ToString()
	b := key1.ToString()
	c := key2.ToString()

	if a != b {
		t.Fatalf("key0 = %s, key1 = %s", a, b)
	}

	if a != c {
		t.Fatalf("key0 = %s, key2 = %s", a, c)
	}
}

func TestWeirdKey(t *testing.T) {
	key0 := NewKey("//a///b/c//")
	expected := NewKey("a/b/c")

	if key0.ToString() != expected.ToString() {
		t.Fail()
	}
}
