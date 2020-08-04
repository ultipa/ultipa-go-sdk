package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ShowPropertyRequest = struct {
	Dataset types.DBType;
}

func (t *Connection) ListProperty (request *types.Request_Property, commonReq *types.Request_Common) *types.ResListProperty {
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
	res := t.UQLListSample(uql.ToString(), commonReq)
	properties := res.Data
	var newData []*types.Response_Property
	for _, pty := range *properties{
		newPty := types.Response_Property{
			Lte: (*pty)["lte"].(string),
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