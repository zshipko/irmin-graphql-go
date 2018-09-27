package irmin

import (
	"context"

	"github.com/shurcooL/graphql"
)

type Info struct {
	Author  graphql.String
	Message graphql.String
}

func (ir Irmin) Set(ctx context.Context, branch string, key string, value string, info *Info) (*Commit, error) {
	type query struct {
		Set struct {
			Commit Commit
		} `graphql:"Set(branch: $branch, key: $key, value: $value, info: $info)"`
	}

	var q query
	vars := map[string]interface{}{
		"branch": graphql.String(branch),
		"key":    graphql.String(key),
		"value":  graphql.String(value),
	}

	if info != nil {
		vars["info"] = info
	}

	err := ir.client.Mutate(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Set.Commit, nil
}
