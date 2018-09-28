package irmin

import (
	"testing"
)

func TestKeyRoundtrip(t *testing.T) {
	key0 := NewKey("/a/b/c")
	key1 := NewKey("a/b/c/")
	key2 := Key{"a", "b", "c"}

	if key0.ToString() != key1.ToString() {
		t.Fatalf("key0 = %s, key1 = %s", key0.ToString(), key1.ToString())
	}

	if key0.ToString() != key2.ToString() {
		t.Fatalf("key0 = %s, key2 = %s", key0.ToString(), key2.ToString())
	}
}
