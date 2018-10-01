package irmin

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

// Commit is used in queries that return commit information
type Commit struct {
	Hash graphql.String
	Info Info
}

// Branch is used in queries that return branch information
type Branch struct {
	Name graphql.String
	Head Commit
}

// ErrNotFound is returned when a key is not available
var ErrNotFound = errors.New("Not found")

// Master - master { ... }
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

// Branch - branch(name: $name)
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

// Commit - commit(hash: $hash)
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

// Get - branch(name: $name) { get(key: $key) }
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
		return []byte{}, ErrNotFound
	}

	return []byte(q.Branch.Get), nil
}
