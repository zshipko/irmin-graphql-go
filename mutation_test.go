package irmin

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestSet(t *testing.T) {
	commit, err := main.Set(context.Background(), NewPath("a/b/c"), []byte("123"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Invalid commit hash")
	}

	info, err := main.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if info.Head.Hash != commit.Hash {
		t.Fatal("Commit hash doesn't match HEAD commit")
	}

	value, err := main.Get(context.Background(), NewPath("a/b/c"))
	if err != nil {
		t.Fatal(err)
	}

	if string(value.Value) != "123" {
		t.Error("Incorrect value after call to set")
	}
}

func TestRemove(t *testing.T) {
	info, err := main.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	commit, err := main.Remove(context.Background(), NewPath("a/b/c"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if info.Head.Hash == commit.Hash {
		t.Fatal("Expected new commit after remove")
	}

	value, err := main.Get(context.Background(), NewPath("a/b/c"))
	if err != ErrNotFound {
		t.Errorf("Expected a/b/c to be removed, instead got '%s'", value)
	}

}

func TestPull(t *testing.T) {
	main.Set(context.Background(), NewPath("README.md"), []byte("something"), nil)

	commit, err := main.Pull(context.Background(), "git://github.com/zshipko/irmin-go", nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit.Hash == "" {
		t.Fatal("Expected new commit hash")
	}

	value, _ := main.Get(context.Background(), NewPath("README.md"))
	readme, _ := ioutil.ReadFile("README.md")
	if value == nil || string(value.Value) != string(readme) {
		t.Fatal("Pull failed")
	}

	// Check contents after pull using List/GetTree

	items, err := main.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(items) == 0 {
		t.Fatal("Expected items")
	}

	if string(items["README.md"]) != string(readme) {
		t.Fatal("Incorrect list results")
	}

	commit, err = main.Set(context.Background(), NewPath("a/b/c"), []byte("123"), nil)
	if err != nil {
		t.Fatal(err)
	}

	x, err := main.GetTree(context.Background(), EmptyPath())
	if err != nil {
		t.Fatal(err)
	}

	if string(x["README.md"].Value) != string(readme) {
		t.Log(x)
		t.Fatal("Incorrect get_tree result for README.md")
	}

	test, _ := ioutil.ReadFile("test.sh")
	if string(x["test.sh"].Value) != string(test) {
		t.Log(x)
		t.Fatal("Incorrect get_tree result for path test.sh")
	}

	if string(x["a/b/c"].Value) != string("123") {
		t.Fatal("Incorrect get_tree result for path a/b/c")
	}
}

func TestSetTree(t *testing.T) {
	path := EmptyPath()
	tree := map[string]Contents{
		"foo/bar/baz": Contents{
			Value:    []byte("testing"),
			Metadata: nil,
		},
	}

	commit, err := main.SetTree(context.Background(), path, tree, nil)
	if err != nil {
		t.Fatal(err)
	}

	if commit == nil {
		t.Fatal("Invalid commit")
	}

	x, err := main.GetTree(context.Background(), path)

	for k, v := range x {
		if string(v.Value) != string(tree[k].Value) {
			t.Fatalf("Invalid value: %s, %s %s", k, string(v.Value), string(tree[k].Value))
		}
	}
}
