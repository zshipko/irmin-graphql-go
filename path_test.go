package irmin

import (
	"testing"
)

func TestPathRoundtrip(t *testing.T) {
	path0 := NewPath("/a/b/c")
	path1 := NewPath("a/b/c/")
	path2 := Path{"a", "b", "c"}

	a := path0.ToString()
	b := path1.ToString()
	c := path2.ToString()

	if a != b {
		t.Fatalf("path0 = %s, path1 = %s", a, b)
	}

	if a != c {
		t.Fatalf("path0 = %s, path2 = %s", a, c)
	}
}

func TestWeirdPath(t *testing.T) {
	path0 := NewPath("//a///b/c//")
	expected := NewPath("a/b/c")

	if path0.ToString() != expected.ToString() {
		t.Fail()
	}
}
