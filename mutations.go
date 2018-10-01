package irmin

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

// Info stores commit information
type Info struct {
	Author  graphql.String
	Message graphql.String
}

// Set a key
func (ir Irmin) Set(ctx context.Context, branch string, key Key, value []byte, info *Info) (*Commit, error) {
	type query struct {
		Set Commit `graphql:"set(branch: $branch, key: $key, value: $value, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"key":    graphql.String(key.ToString()),
		"value":  graphql.String(value),
		"info":   info,
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Set, nil
}

// Remove a key
func (ir Irmin) Remove(ctx context.Context, branch string, key Key, info *Info) (*Commit, error) {
	type query struct {
		Remove Commit `graphql:"remove(branch: $branch, key: $key, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"key":    graphql.String(key.ToString()),
		"info":   info,
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Remove, nil
}

// Merge two branches
func (ir Irmin) Merge(ctx context.Context, branch string, fromBranch string, info *Info) (*Commit, error) {
	type query struct {
		Merge Commit `graphql:"merge(branch: $branch, from: $from, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"from":   fromBranch,
		"info":   info,
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Merge, nil
}

// Revert to the given snapshot
func (ir Irmin) Revert(ctx context.Context, branch string, hash string) (*Commit, error) {
	type query struct {
		Revert Commit `graphql:"revert(branch: $branch, commit: $commit)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"commit": graphql.String(hash),
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Revert, nil
}

// Pull from a remote store
func (ir Irmin) Pull(ctx context.Context, branch string, remote string) (*Commit, error) {
	type query struct {
		Pull Commit `graphql:"pull(branch: $branch, remote: $remote)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"remote": graphql.String(remote),
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Pull, nil
}

// Push to a remote store
func (ir Irmin) Push(ctx context.Context, branch string, remote string) error {
	type query struct {
		Push graphql.String `graphql:"pull(branch: $branch, remote: $remote)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"remote": graphql.String(remote),
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return err
	}

	if q.Push != "" {
		return errors.New(string(q.Push))
	}

	return nil
}
