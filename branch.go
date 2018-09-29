package irmin

import (
	"context"
)

type BranchRef struct {
	client *Irmin
	name   string
}

func (ir *Irmin) GetBranch(name string) BranchRef {
	return BranchRef{
		client: ir,
		name:   name,
	}
}

func (b BranchRef) Get(ctx context.Context, key Key) ([]byte, error) {
	return b.client.Get(ctx, b.name, key)
}

func (b BranchRef) Set(ctx context.Context, key Key, value []byte, info *Info) (*Commit, error) {
	return b.client.Set(ctx, b.name, key, value, info)
}

func (b BranchRef) Pull(ctx context.Context, remote string) (*Commit, error) {
	return b.client.Pull(ctx, b.name, remote)
}
