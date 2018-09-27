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
    "github.com/zshipko/irmin-go"
)

func main(){
    client := irmin.New("http://localhost:1234/graphql", nil)
    if client == nil {
        panic("Unable to connect to irmin-graphql server")
    }

    master := client.Master()
    fmt.Println(master.Head.Hash)
}
```
