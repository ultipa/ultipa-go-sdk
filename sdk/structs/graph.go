package structs

import "ultipa-go-sdk/sdk/types"

type Graph struct {
	ID types.ID
	Name string
	Description string
	TotalNodes uint64
	TotalEdges uint64
}

