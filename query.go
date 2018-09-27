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
func (ir Irmin) Master(ctx context.Context) (*Branch, error) {
	type query struct {
		Master Branch
	}

	var q query
	err := ir.client.Query(ctx, &q, nil)
	if err != nil {
		return nil, err
	}

	return &q.Master, nil
}

// branch(name: $name)
func (ir Irmin) Branch(ctx context.Context, name string) (*Branch, error) {
	type query struct {
		Branch Branch `graphql:"branch(name: $name)"`
	}

	var q query
	vars := map[string]interface{}{
		"name": graphql.String(name),
	}
	err := ir.client.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Branch, nil
}

// commit(hash: $hash)
func (ir Irmin) Commit(ctx context.Context, hash string) (*Commit, error) {
	type query struct {
		Commit Commit `graphql:"commit(hash: $hash)"`
	}

	var q query
	vars := map[string]interface{}{
		"hash": graphql.String(hash),
	}
	err := ir.client.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Commit, nil
}
