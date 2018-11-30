package irmin

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestSet(t *testing.T) {
	commit, err := master.Set(context.Background(), NewKey("a/b/c"), []byte("123"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Invalid commit hash")
	}

	info, err := master.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if info.Head.Hash != commit.Hash {
		t.Fatal("Commit hash doesn't match HEAD commit")
	}

	value, err := master.Get(context.Background(), NewKey("a/b/c"))
	if err != nil {
		t.Fatal(err)
	}

	if string(value) != "123" {
		t.Error("Incorrect value after call to set")
	}
}

func TestRemove(t *testing.T) {
	info, err := master.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	commit, err := master.Remove(context.Background(), NewKey("a/b/c"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if info.Head.Hash == commit.Hash {
		t.Fatal("Expected new commit after remove")
	}

	value, err := master.Get(context.Background(), NewKey("a/b/c"))
	if err != ErrNotFound {
		t.Errorf("Expected a/b/c to be removed, instead got '%s'", value)
	}

}

func TestPull(t *testing.T) {
	master.Set(context.Background(), NewKey("README.md"), []byte("something"), nil)

	commit, err := master.Pull(context.Background(), "git://github.com/zshipko/irmin-go.git")
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Expected new commit hash")
	}

	value, _ := master.Get(context.Background(), NewKey("README.md"))
	readme, _ := ioutil.ReadFile("README.md")
	if string(value) != string(readme) {
		t.Fatal("Pull failed")
	}

	// Check contents after pull using List/GetTree

	items, err := master.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(items) == 0 {
		t.Fatal("Expected items")
	}

	if string(items["README.md"]) != string(readme) {
		t.Log(items)
		t.Log(string(items["README.md"]))
		t.Fatal("Incorrect list results")
	}

	commit, err = master.Set(context.Background(), NewKey("a/b/c"), []byte("123"), nil)
	if err != nil {
		t.Fatal(err)
	}

	x, err := master.GetTree(context.Background(), EmptyKey())

	if string(x["README.md"].Value) != string(readme) {
		t.Log(x)
		t.Fatal("Incorrect get_tree result for README.md")
	}

	test, _ := ioutil.ReadFile("test.sh")
	if string(x["test.sh"].Value) != string(test) {
		t.Log(x)
		t.Fatal("Incorrect get_tree result for key test.sh")
	}

	if string(x["a/b/c"].Value) != string("123") {
		t.Fatal("Incorrect get_tree result for key a/b/c")
	}
}

func TestSetTree(t *testing.T) {
	key := NewKey("/foo")
	tree := map[string]Contents{
		"bar/baz": Contents{
			Value:    []byte("testing"),
			Metadata: nil,
		},
	}

	commit, err := master.SetTree(context.Background(), key, tree, nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit == nil {
		t.Fatal("Invalid commit")
	}

	x, err := master.GetTree(context.Background(), key)

	for k, v := range x {
		if string(v.Value) != string(tree[k].Value) {
			t.Fatalf("Invalid value: %s", k)
		}
	}
}
