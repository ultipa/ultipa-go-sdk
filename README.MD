# Ultipa GO SDK

## Installation

```shell

go get github.com/ultipa/ultipa-go-sdk
go mod tidy

```

## Document

Read Ultipa Go SDK document  [ultipa.com](https://www.ultipa.com/document/ultipa-drivers/go-installation/v4.3)

## Create Client

```go

import (
    "github.com/ultipa/ultipa-go-sdk/sdk"
    "github.com/ultipa/ultipa-go-sdk/sdk/api"
    "github.com/ultipa/ultipa-go-sdk/sdk/configuration"
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

- Alias(alias string) : *DataItem
- Alias(index int) : *DataItem


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

---


## Trouble Shooting
### grpc issues
This sdk depends on Protocol Buffer and grpc, so we need to compile the proto file first:<br>
go install google.golang.org/grpc/cmd/protoc-gen-go@latest<br>
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest<br>

then execute the build shell:<br>
cd ./rpc <br>
sh ./build-proto.sh

after this, remove protoc-gen-go and protoc-gen-go-grpc dependencies from go.mod



