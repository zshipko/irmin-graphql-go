package irmin

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestSet(t *testing.T) {
	commit, err := client.Set(context.Background(), "master", NewKey("a/b/c"), []byte("123"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Invalid commit hash")
	}

	master, err := client.Branch(context.Background(), "master")
	if err != nil {
		t.Fatal(err)
	}

	if master.Head.Hash != commit.Hash {
		t.Fatal("Commit hash doesn't match HEAD commit")
	}

	value, err := client.Get(context.Background(), "master", NewKey("a/b/c"))
	if err != nil {
		t.Fatal(err)
	}

	if string(value) != "123" {
		t.Error("Incorrect value after call to set")
	}
}

func TestRemove(t *testing.T) {
	master, err := client.Branch(context.Background(), "master")
	if err != nil {
		t.Fatal(err)
	}

	commit, err := client.Remove(context.Background(), "master", NewKey("a/b/c"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if master.Head.Hash == commit.Hash {
		t.Fatal("Expected new commit after remove")
	}

	value, err := client.Get(context.Background(), "master", NewKey("a/b/c"))
	if err != NotFound {
		t.Errorf("Expected a/b/c to be removed, instead got '%s'", value)
	}

}

func TestPull(t *testing.T) {
	client.Set(context.Background(), "master", NewKey("README.md"), []byte("something"), nil)

	commit, err := client.Pull(context.Background(), "master", "git://github.com/zshipko/irmin-go.git")
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Expected new commit hash")
	}

	value, _ := client.Get(context.Background(), "master", NewKey("README.md"))
	readme, _ := ioutil.ReadFile("README.md")
	if string(value) != string(readme) {
		t.Fatal("Pull failed")
	}
}
