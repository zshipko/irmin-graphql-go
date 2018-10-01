package irmin

import (
	"context"
)

// BranchRef is a reference to the branch specified by `name`
type BranchRef struct {
	client *Irmin
	name   string
}

// GetBranch returns a BranchRef
func (ir *Irmin) GetBranch(name string) BranchRef {
	return BranchRef{
		client: ir,
		name:   name,
	}
}

// Get a key
func (b BranchRef) Get(ctx context.Context, key Key) ([]byte, error) {
	return b.client.Get(ctx, b.name, key)
}

// Set a key
func (b BranchRef) Set(ctx context.Context, key Key, value []byte, info *Info) (*Commit, error) {
	return b.client.Set(ctx, b.name, key, value, info)
}

// Pull from a remote store
func (b BranchRef) Pull(ctx context.Context, remote string) (*Commit, error) {
	return b.client.Pull(ctx, b.name, remote)
}
