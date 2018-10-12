package irmin

import (
	"strings"
)

// Key attempts to mimic Irmin keys
type Key []string

// ToString converts a key to a string
func (k Key) ToString() *string {
	if len(k) == 0 {
		return nil
	}

	s := strings.Join(k, "/")
	return &s
}

// NewKey creates a new key from a string
func NewKey(key string) Key {
	key = strings.Trim(key, "/")
	return strings.Split(key, "/")
}

// EmptyKey create a new empty key
func EmptyKey() Key {
	return []string{}
}
