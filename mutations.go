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
func (br BranchRef) Set(ctx context.Context, key Key, value []byte, info *Info) (*Commit, error) {
	type query struct {
		Set Commit `graphql:"set(branch: $branch, key: $key, value: $value, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"key":   key.Arg(),
		"value": graphql.String(value),
		"info":  info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Set, nil
}

// SetAll allows you to set a key/value pair with metadata
func (br BranchRef) SetAll(ctx context.Context, key Key, value []byte, metadata []byte, info *Info) (*Commit, error) {
	type query struct {
		SetAll Commit `graphql:"set(branch: $branch, key: $key, value: $value, metadata: $metadata, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"key":      key.Arg(),
		"value":    graphql.String(value),
		"metadata": graphql.String(metadata),
		"info":     info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.SetAll, nil
}

// Remove a key
func (br BranchRef) Remove(ctx context.Context, key Key, info *Info) (*Commit, error) {
	type query struct {
		Remove Commit `graphql:"remove(branch: $branch, key: $key, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"key":  key.Arg(),
		"info": info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Remove, nil
}

// Merge two branches
func (br BranchRef) Merge(ctx context.Context, fromBranch string, info *Info) (*Commit, error) {
	type query struct {
		Merge Commit `graphql:"merge(branch: $branch, from: $from, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"from": fromBranch,
		"info": info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Merge, nil
}

// Revert to the given snapshot
func (br BranchRef) Revert(ctx context.Context, hash string) (*Commit, error) {
	type query struct {
		Revert Commit `graphql:"revert(branch: $branch, commit: $commit)"`
	}

	var q query
	vars := map[string]interface{}{
		"commit": graphql.String(hash),
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Revert, nil
}

// Pull from a remote store
func (br BranchRef) Pull(ctx context.Context, remote string) (*Commit, error) {
	type query struct {
		Pull Commit `graphql:"pull(branch: $branch, remote: $remote)"`
	}

	var q query
	vars := map[string]interface{}{
		"remote": graphql.String(remote),
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Pull, nil
}

// Push to a remote store
func (br BranchRef) Push(ctx context.Context, remote string) error {
	type query struct {
		Push graphql.String `graphql:"pull(branch: $branch, remote: $remote)"`
	}

	var q query
	vars := map[string]interface{}{
		"remote": graphql.String(remote),
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return err
	}

	if q.Push != "" {
		return errors.New(string(q.Push))
	}

	return nil
}
