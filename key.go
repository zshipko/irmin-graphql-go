package irmin

import (
	"strings"
)

// Key attempts to mimic Irmin keys
type Key []string

// ToString converts a key to a string
func (k Key) ToString() string {
	return strings.Join(k, "/")
}

// NewKey creates a new key from a string
func NewKey(key string) Key {
	key = strings.Trim(key, "/")
	return strings.Split(key, "/")
}
