package irmin

import (
	"context"
	"testing"
)

func TestSet(t *testing.T) {
	commit, err := client.Set(context.Background(), "master", NewKey("a/b/c"), "123", nil)
	if err != nil {
		t.Error(err)
	}

	if commit.Hash == "" {
		t.Error("Invalid commit hash")
	}
}
