package irmin

import (
	"context"

	"github.com/shurcooL/graphql"
)

// BranchRef is a reference to the branch specified by `name`
type BranchRef struct {
	*Irmin
	name string
}

// Query sends a query to the specified graphql server
func (b BranchRef) Query(ctx context.Context, q interface{}, vars map[string]interface{}) error {
	if vars == nil {
		vars = map[string]interface{}{}
	}

	vars["branch"] = graphql.String(b.name)
	return b.Client.Query(ctx, q, vars)
}

// Mutate sends a mutation to the specified graphql server
func (b BranchRef) Mutate(ctx context.Context, q interface{}, vars map[string]interface{}) error {
	if vars == nil {
		vars = map[string]interface{}{}
	}

	vars["branch"] = graphql.String(b.name)
	return b.Client.Mutate(ctx, q, vars)
}
