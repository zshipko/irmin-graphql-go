package irmin

import (
	"context"
	"testing"
)

var client *Irmin = nil

func init() {
	client = New("http://127.0.0.1:8080/graphql", nil)
	if client == nil {
		panic("Invalid client")
	}
}

func TestMaster(t *testing.T) {
	master, err := client.Master(context.Background())
	if err != nil {
		t.Error(err)
	}

	if master.Name != "master" {
		t.Error("Expected master branch")
	}
}
