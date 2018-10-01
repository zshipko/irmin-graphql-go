package irmin

import (
	"github.com/shurcooL/graphql"
	"net/http"
)

// Irmin client
type Irmin struct {
	client *graphql.Client
}

// New creates a new Irmin client
func New(url string, conn *http.Client) *Irmin {
	client := graphql.NewClient(url, conn)
	if client == nil {
		return nil
	}

	return &Irmin{
		client: client,
	}
}
