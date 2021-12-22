# Ultipa GO SDK(Driver 4.0)

---

## Create Client

```go

import (
    "ultipa-go-sdk/sdk"
    "ultipa-go-sdk/sdk/api"
    "ultipa-go-sdk/sdk/configuration"
)

config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
    Hosts: []string{"127.0.0.1:60061"},
    Username: "root",
    Password: "root",
    DefaultGraph: "default",
})

client, err = sdk.NewUltipa(config)
```

## Send UQL and get result

```go
resp, err := client.UQL("find().nodes() return nodes limit 1", nil)
nodes := resp.alias("nodes").asNodes()
```

## UQLResponse Interface

- Alias(alias string) : DataItem
- Alias(index int) : DataItem


## DataItem Interface

- asNodes() : (nodes []*structs.Node, schemas map[string]*structs.Schema, err error)
- asEdges() : (edges []*structs.Edge, schemas map[string]*structs.Schema, err error)
- asPaths() : (paths []*structs.Path, err error)
- asTable() : (table *structs.Table, err error)
- asArray() : (arr *structs.Array, err error)
- asAttr()  : (arr *structs.Array, err error)
- asGraphs() : (graphs []*structs.Graph, err error)
- asSchemas() : (schemas []*structs.Schema, err error)
- asAlgos() : ([]*structs.Algo, error)
- asAny() : (interface{}, error)

## Batch Insert

Check Go Transporter For Details




