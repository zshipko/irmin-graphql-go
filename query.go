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

// TreeMap is used in GetTree and SetTree to represent a listing of tree nodes
type TreeMap map[string]Contents

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

// GetTree - branch(name: $branch) { tree { get_tree(key: $key) { list_contents_recursively { key, value, metadata } } } }
func (br BranchRef) GetTree(ctx context.Context, key Key) (TreeMap, error) {
	type query struct {
		Branch struct {
			Tree struct {
				GetTree struct {
					List []struct {
						Key      graphql.String `graphql:"key"`
						Value    graphql.String `graphql:"value"`
						Metadata graphql.String `graphql:"metadata"`
					} `graphql:"list_contents_recursively"`
				} `graphql:"get_tree(key: $key)"`
			} `graphql:"tree"`
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

	items := map[string]Contents{}

	for _, v := range q.Branch.Tree.GetTree.List {
		items[NewKey(string(v.Key)).ToString()] = Contents{
			Value:    []byte(v.Value),
			Metadata: []byte(v.Metadata),
		}
	}

	return items, nil
}

// Get - branch(name: $branch) { tree { get_contents(key: $key) { value, metadata } } }
func (br BranchRef) Get(ctx context.Context, key Key) (*Contents, error) {
	type query struct {
		Branch struct {
			Tree struct {
				Get *struct {
					Value    graphql.String `graphql:"value"`
					Metadata graphql.String `graphql:"metadata"`
				} `graphql:"get_contents(key: $key)"`
			} `graphql:"tree"`
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

	if q.Branch.Tree.Get == nil {
		return nil, ErrNotFound
	}

	return &Contents{
		Value:    []byte(q.Branch.Tree.Get.Value),
		Metadata: []byte(q.Branch.Tree.Get.Metadata),
	}, nil
}

// List returns a list of the values stored under the specified key
func (br BranchRef) List(ctx context.Context, key Key) (map[string][]byte, error) {
	type query struct {
		Branch struct {
			Head struct {
				Tree struct {
					GetTree struct {
						List []struct {
							Key   graphql.String
							Value graphql.String
						} `graphql:"list"`
					} `graphql:"get_tree(key: $key)"`
				} `graphql:"tree"`
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
	for _, x := range q.Branch.Head.Tree.GetTree.List {
		items[NewKey(string(x.Key)).ToString()] = []byte(x.Value)
	}

	return items, nil
}
