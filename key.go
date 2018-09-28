package irmin

import (
	"strings"
)

type Key []string

func (k Key) ToString() string {
	return strings.Join(k, "/")
}

func NewKey(key string) Key {
	return strings.Split(key, "/")
}
