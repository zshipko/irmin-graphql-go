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

// ErrNotFound is returned when a path is not available
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

// GetTree - branch(name: $branch) { tree { get_tree(path: $path) { list_contents_recursively { path, value, metadata } } } }
func (br BranchRef) GetTree(ctx context.Context, path Path) (TreeMap, error) {
	type query struct {
		Branch struct {
			Tree struct {
				GetTree struct {
					List []struct {
						Path      graphql.String `graphql:"path"`
						Value    graphql.String `graphql:"value"`
						Metadata graphql.String `graphql:"metadata"`
					} `graphql:"list_contents_recursively"`
				} `graphql:"get_tree(path: $path)"`
			} `graphql:"tree"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"path": path.Arg(),
	}

	err := br.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	items := map[string]Contents{}

	for _, v := range q.Branch.Tree.GetTree.List {
		items[NewPath(string(v.Path)).ToString()] = Contents{
			Value:    []byte(v.Value),
			Metadata: []byte(v.Metadata),
		}
	}

	return items, nil
}

// Get - branch(name: $branch) { tree { get_contents(path: $path) { value, metadata } } }
func (br BranchRef) Get(ctx context.Context, path Path) (*Contents, error) {
	type query struct {
		Branch struct {
			Tree struct {
				Get *struct {
					Value    graphql.String `graphql:"value"`
					Metadata graphql.String `graphql:"metadata"`
				} `graphql:"get_contents(path: $path)"`
			} `graphql:"tree"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"path": path.Arg(),
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

// List returns a list of the values stored under the specified path
func (br BranchRef) List(ctx context.Context, path Path) (map[string][]byte, error) {
	type query struct {
		Branch struct {
			Head struct {
				Tree struct {
					GetTree struct {
						List []struct {
							Path   graphql.String
							Value graphql.String
						} `graphql:"list"`
					} `graphql:"get_tree(path: $path)"`
				} `graphql:"tree"`
			} `graphql:"head"`
		} `graphql:"branch(name: $branch)"`
	}

	var q query
	vars := map[string]interface{}{
		"path": path.Arg(),
	}

	err := br.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	items := map[string][]byte{}
	for _, x := range q.Branch.Head.Tree.GetTree.List {
		items[NewPath(string(x.Path)).ToString()] = []byte(x.Value)
	}

	return items, nil
}
