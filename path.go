package irmin

import (
	"strings"

	"github.com/shurcooL/graphql"
)

// Path attempts to mimic Irmin paths
type Path []string

// ToString converts a path to a string
func (k Path) ToString() string {
	if len(k) == 0 {
		return ""
	}

	s := strings.Join(k, "/")
	return s
}

// Arg converts a path to a GraphQL query argument
func (k Path) Arg() *graphql.String {
	return graphql.NewString(graphql.String(k.ToString()))
}

// NewPath creates a new path from a string
func NewPath(path string) Path {
	path = strings.Trim(path, "/")
	tmp := strings.Split(path, "/")
	dest := []string{}

	for _, c := range tmp {
		if c != "" {
			dest = append(dest, c)

		}
	}

	return dest
}

// EmptyPath create a new empty path
func EmptyPath() Path {
	return []string{}
}
