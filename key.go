package irmin

import (
	"strings"

	"github.com/shurcooL/graphql"
)

// Key attempts to mimic Irmin keys
type Key []string

// ToString converts a key to a string
func (k Key) ToString() string {
	if len(k) == 0 {
		return ""
	}

	s := strings.Join(k, "/")
	return s
}

// Arg converts a key to a GraphQL query argument
func (k Key) Arg() *graphql.String {
	return graphql.NewString(graphql.String(k.ToString()))
}

// NewKey creates a new key from a string
func NewKey(key string) Key {
	key = strings.Trim(key, "/")
	tmp := strings.Split(key, "/")
	dest := []string{}

	for _, c := range tmp {
		if c != "" {
			dest = append(dest, c)

		}
	}

	return dest
}

// EmptyKey create a new empty key
func EmptyKey() Key {
	return []string{}
}
