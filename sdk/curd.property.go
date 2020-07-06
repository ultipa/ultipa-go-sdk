package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/types/types_response"
	"ultipa-go-sdk/utils"
)

type ShowPropertyRequest = struct {
	Dataset types.DBType;
}

func (t *Connection) ListProperty (request ShowPropertyRequest) *types.ResListProperty {
	uql := utils.UQLMAKER{}
	dataset := request.Dataset
	switch dataset {
	case types.DBType_DBNODE:
		uql.SetCommand(utils.CommandList_showNodeProperty)
		break
	case types.DBType_DBEDGE:
		uql.SetCommand(utils.CommandList_showEdgeProperty)
		break
	}
	res := t.UQL(uql.ToString(), nil)
	uqlReply := res.Data
	properties := utils.TableToArray((*uqlReply.Tables)[0])
	var newData []*types_response.Property
	for _, pty := range *properties{
		newPty := types_response.Property{
			Lte: (*pty)["lte"] == "true",
			Index: (*pty)["index"] == "true",
			PropertyName: (*pty)["name"].(string),
			PropertyType: (*pty)["type"].(string),
		}
		newData = append(newData, &newPty)
	}


	return &types.ResListProperty{
		res.ResWithoutData,
		newData,
	}
}