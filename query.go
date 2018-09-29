package irmin

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

type Commit struct {
	Hash graphql.String
}

type Branch struct {
	Name graphql.String
	Head Commit
}

var NotFound error = errors.New("Not found")

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

// branch(name: $name) { get(key: $key) }
func (ir Irmin) Get(ctx context.Context, branch string, key Key) ([]byte, error) {
	type query struct {
		Branch struct {
			Get graphql.String `graphql:"get(key: $key)"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"key":    graphql.String(key.ToString()),
	}

	err := ir.client.Query(ctx, &q, vars)
	if err != nil {
		return []byte{}, err
	}

	if len(q.Branch.Get) == 0 {
		return []byte{}, NotFound
	}

	return []byte(q.Branch.Get), nil
}
