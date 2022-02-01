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

// TreeItem is used to encode tree contents int GraphQL queries
type TreeItem map[string]interface{}

func makeTreeArray(tree TreeMap) []TreeItem {
	treeArray := []TreeItem{}

	for k, v := range tree {
		var value *graphql.String
		if v.Value != nil {
			value = graphql.NewString(graphql.String(string(v.Value)))
		}

		var meta *graphql.String
		if v.Metadata != nil {
			meta = graphql.NewString(graphql.String(string(v.Metadata)))
		}

		treeArray = append(treeArray, TreeItem{
			"path":      graphql.String(k),
			"value":    value,
			"metadata": meta,
		})
	}

	return treeArray
}

// Set a path
func (br BranchRef) Set(ctx context.Context, path Path, value []byte, info *Info) (*Commit, error) {
	type query struct {
		Set Commit `graphql:"set(branch: $branch, path: $path, value: $value, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"path":   path.Arg(),
		"value": graphql.String(value),
		"info":  info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Set, nil
}

// UpdateTree allows you to modify multiple paths at the same time
func (br BranchRef) UpdateTree(ctx context.Context, path Path, tree TreeMap, info *Info) (*Commit, error) {
	type query struct {
		UpdateTree Commit `graphql:"update_tree(branch: $branch, path: $path, tree: $tree, info: $info)"`
	}

	treeArray := makeTreeArray(tree)

	var q query
	vars := map[string]interface{}{
		"path":  path.Arg(),
		"tree": treeArray,
		"info": info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.UpdateTree, nil
}

// SetTree allows you to set multiple paths at the same time
func (br BranchRef) SetTree(ctx context.Context, path Path, tree TreeMap, info *Info) (*Commit, error) {
	type query struct {
		SetTree Commit `graphql:"set_tree(branch: $branch, path: $path, tree: $tree, info: $info)"`
	}

	treeArray := makeTreeArray(tree)

	var q query
	vars := map[string]interface{}{
		"path":  path.Arg(),
		"tree": treeArray,
		"info": info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.SetTree, nil
}

// SetAll allows you to set a path/value pair with metadata
func (br BranchRef) SetAll(ctx context.Context, path Path, value []byte, metadata []byte, info *Info) (*Commit, error) {
	type query struct {
		SetAll Commit `graphql:"set(branch: $branch, path: $path, value: $value, metadata: $metadata, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"path":      path.Arg(),
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

// Remove a path
func (br BranchRef) Remove(ctx context.Context, path Path, info *Info) (*Commit, error) {
	type query struct {
		Remove Commit `graphql:"remove(branch: $branch, path: $path, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"path":  path.Arg(),
		"info": info,
	}

	err := br.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Remove, nil
}

// MergeWithBranch merges two branches
func (br BranchRef) MergeWithBranch(ctx context.Context, fromBranch string, info *Info) (*Commit, error) {
	type query struct {
		Merge Commit `graphql:"merge_with_branch(branch: $branch, from: $from, info: $info)"`
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
func (br BranchRef) Pull(ctx context.Context, remote string, info *Info) (*Commit, error) {
	type query struct {
		Pull Commit `graphql:"pull(branch: $branch, remote: $remote, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"remote": graphql.String(remote),
		"info":   info,
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
