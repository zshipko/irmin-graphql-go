# irmin-graphql-go

Go bindings to [Irmin](https://github.com/mirage/irmin) using GraphQL.

## Installation

```shell
$ go get -u github.com/zshipko/irmin-graphql-go
```

## Getting started

The following example will print the commit hash of the current `HEAD` of the main branch:

```go
package main

import (
    "fmt"
    "context"

    "github.com/zshipko/irmin-graphql-go"
)

func main(){
    client := irmin.New("http://localhost:8080/graphql", nil)
    if client == nil {
        panic("Unable to connect to irmin-graphql server")
    }

    main := client.Main(context.Background())
    fmt.Println(main.Head.Hash)
}
```

## Writing custom queries

`irmin-go` provides many predefined queries, but at some point you may need to define your own query. Thanks to [github.com/shurcooL/graphql](https://github.com/shurcooL/graphql) this is very simple:

```go
package main

import (
	"context"
	"log"

	"github.com/shurcooL/graphql"
	"github.com/zshipko/irmin-go"
)

func main() {
	client := irmin.New("http://localhost:8080/graphql", nil)
	if client == nil {
		panic("Unable to connect to irmin-graphql server")
	}

	var query struct {
	  Main struct {
      Tree struct {
			  Get *graphql.String `graphql:"get(key: $key)"`
      }
		}
	}

	keys := map[string]interface{}{
		"key": "a/b/c",
	}

	err := client.Client.Query(context.Background(), &query, keys)
	if err != nil {
		log.Fatal(err)
	}

	if query.Main.Tree.Get == nil {
		log.Println("NULL")
	} else {
		log.Println(*query.Main.Tree.Get)
	}
}
```
