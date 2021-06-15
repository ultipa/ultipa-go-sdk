package structs

import "ultipa-go-sdk/sdk/types"

type MetaData struct {
	ID     types.ID
	UUID   types.UUID
	From   types.ID
	To     types.ID
	Schema types.Schema
	Values types.Values
}


