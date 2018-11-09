# irmin-go

Go bindings to [Irmin](https://github.com/mirage/irmin) using GraphQL.

## Installation

```shell
$ go get -u github.com/zshipko/irmin-go
```

## Getting started

The following example will print the commit hash of the current `HEAD` of the master branch:

```go
package main

import (
    "fmt"
    "context"

    "github.com/zshipko/irmin-go"
)

func main(){
    client := irmin.New("http://localhost:8080/graphql", nil)
    if client == nil {
        panic("Unable to connect to irmin-graphql server")
    }

    master := client.Master(context.Background())
    fmt.Println(master.Head.Hash)
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
		Master struct {
			Get *graphql.String `graphql:"get(key: $key)"`
		}
	}

	keys := map[string]interface{}{
		"key": "a/b/c",
	}

	err := client.Client.Query(context.Background(), &query, keys)
	if err != nil {
		log.Fatal(err)
	}

	if query.Master.Get == nil {
		log.Println("NULL")
	} else {
		log.Println(*query.Master.Get)
	}
}
```
