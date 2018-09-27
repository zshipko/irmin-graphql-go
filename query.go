package irmin

import (
	"context"
	"github.com/shurcooL/graphql"
)

type Commit struct {
	Hash graphql.String
}

type Branch struct {
	Name graphql.String
	Head Commit
}

// master { ... }
func (ir Irmin) Master() (*Branch, error) {
	type query struct {
		master Branch
	}

	var q query
	err := ir.client.Query(context.Background(), &q, nil)
	if err != nil {
		return nil, err
	}

	return &q.master, nil
}

// branch(name: $name)
func (ir Irmin) Branch(name string) (*Branch, error) {
	type query struct {
		branch Branch `graphql:"branch(name: $name)"`
	}

	var q query
	vars := map[string]interface{}{
		"name": graphql.String(name),
	}
	err := ir.client.Query(context.Background(), &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.branch, nil
}
