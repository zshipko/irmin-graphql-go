package irmin

import (
	"context"

	"github.com/shurcooL/graphql"
)

type Info struct {
	Author  graphql.String
	Message graphql.String
}

func (ir Irmin) Set(ctx context.Context, branch string, key Key, value string, info *Info) (*Commit, error) {
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
