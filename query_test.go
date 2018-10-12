package irmin

import (
	"context"
	"testing"
)

var client *Irmin
var master BranchRef

func init() {
	client = New("http://127.0.0.1:8080/graphql", nil)
	if client == nil {
		panic("Invalid client")
	}
	master = client.Master()
}

func TestMaster(t *testing.T) {
	info, err := master.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if string(info.Name) != "master" {
		t.Error("Expected master branch")
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
