package irmin

import (
	"context"
	"testing"
)

var client *Irmin
var main BranchRef

func init() {
	client = New("http://127.0.0.1:8080/graphql", nil)
	if client == nil {
		panic("Invalid client")
	}
	main = client.Main()
}

func TestMain(t *testing.T) {
	info, err := main.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if string(info.Name) != "main" {
		t.Error("Expected main branch")
	}
}

func TestBranch(t *testing.T) {
	testing, err := client.Branch("testing").Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if testing.Name != "testing" {
		t.Error("Expected testing branch")
	}
}
