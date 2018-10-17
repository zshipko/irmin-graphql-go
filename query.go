package irmin

import (
	"context"
	"errors"

	"github.com/shurcooL/graphql"
)

// Commit is used in queries that return commit information
type Commit struct {
	Hash string
	Info Info
}

// Branch is used in queries that return branch information
type Branch struct {
	Name string
	Head Commit
}

// Contents is used to store contents + metadata
type Contents struct {
	Value    []byte
	Metadata []byte
}

// ErrNotFound is returned when a key is not available
var ErrNotFound = errors.New("Not found")

// Info - branch(name: $branch)
func (br BranchRef) Info(ctx context.Context) (*Branch, error) {
	type query struct {
		Branch Branch `graphql:"branch(name: $branch)"`
	}

	var q query
	err := br.Query(ctx, &q, nil)
	if err != nil {
		return nil, err
	}

	return &q.Branch, nil
}

// Get - branch(name: $branch) { get(key: $key) }
func (br BranchRef) Get(ctx context.Context, key Key) ([]byte, error) {
	type query struct {
		Branch struct {
			Get *graphql.String `graphql:"get(key: $key)"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"key": key.Arg(),
	}

	err := br.Query(ctx, &q, vars)
	if err != nil {
		return []byte{}, err
	}

	if q.Branch.Get == nil {
		return []byte{}, ErrNotFound
	}

	return []byte(*q.Branch.Get), nil
}

// GetAll - branch(name: $branch) { get_all(key: $key) }
func (br BranchRef) GetAll(ctx context.Context, key Key) (*Contents, error) {
	type query struct {
		Branch struct {
			GetAll *struct {
				Value    graphql.String
				Metadata graphql.String
			} `graphql:"get_all(key: $key)"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"key": key.Arg(),
	}

	err := br.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	if q.Branch.GetAll == nil {
		return nil, ErrNotFound
	}

	return &Contents{
		Value:    []byte(q.Branch.GetAll.Value),
		Metadata: []byte(q.Branch.GetAll.Metadata),
	}, nil
}

// List returns a list of the values stored under the specified key
func (br BranchRef) List(ctx context.Context, key Key) (map[string][]byte, error) {
	type query struct {
		Branch struct {
			Head struct {
				Node struct {
					Get struct {
						Tree []struct {
							Key   graphql.String
							Value graphql.String
						} `graphql:"tree"`
					} `graphql:"get(key: $key)"`
				} `graphql:"node"`
			} `graphql:"head"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"key": key.Arg(),
	}

	err := br.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	items := map[string][]byte{}
	for _, x := range q.Branch.Head.Node.Get.Tree {
		items[NewKey(string(x.Key)).ToString()] = []byte(x.Value)
	}

	return items, nil
}
