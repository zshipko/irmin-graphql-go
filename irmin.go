package irmin

import (
	"context"
	"net/http"

	"github.com/shurcooL/graphql"
)

// Irmin client
type Irmin struct {
	Client *graphql.Client
}

// New creates a new Irmin client
func New(url string, conn *http.Client) *Irmin {
	client := graphql.NewClient(url, conn)
	if client == nil {
		return nil
	}

	return &Irmin{
		Client: client,
	}
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
	err := ir.Client.Query(ctx, &q, vars)
	if err != nil {
		return nil, err
	}

	return &q.Commit, nil
}

// Branches returns with a list of available branches
func (ir *Irmin) Branches(ctx context.Context) ([]string, error) {
	type query struct {
		Branches []string
	}

	var q query
	err := ir.Client.Query(ctx, &q, nil)
	if err != nil {
		return []string{}, err
	}

	return q.Branches, nil
}

// Branch returns a BranchRef, which is used to send queries/mutations
func (ir *Irmin) Branch(name string) BranchRef {
	return BranchRef{
		Irmin: ir,
		name:  name,
	}
}

// Master returns a BranchRef for the master branch
func (ir *Irmin) Master() BranchRef {
	return BranchRef{
		Irmin: ir,
		name:  "master",
	}
}
