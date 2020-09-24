package sdk

import (
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ShowPropertyRequest = struct {
	Dataset types.DBType
}

func (t *Connection) ListProperty(request *types.Request_Property, commonReq *types.Request_Common) *types.ResListProperty {
	uql := utils.UQLMAKER{}
	dataset := request.Dataset
	switch dataset {
	case types.DBType_DBNODE:
		uql.SetCommand(utils.UQLCommand_showNodeProperty)
		break
	case types.DBType_DBEDGE:
		uql.SetCommand(utils.UQLCommand_showEdgeProperty)
		break
	}
	res := t.UQLListSample(uql.ToString(), commonReq)
	properties := res.Data
	var newData []*types.Response_Property
	for _, pty := range *properties {
		newPty := types.Response_Property{
			Lte:          (*pty)["lte"].(string),
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

type ResponseCreateProperty struct {
	*types.ResWithoutData
}

func (t *Connection) CreateProperty(dataset ultipa.DBType, propertyName string, propertyType types.PropertyTypeString, description string, commonReq *types.Request_Common) *ResponseCreateProperty {
	uql := utils.UQLMAKER{}
	switch dataset {
	case types.DBType_DBNODE:
		uql.SetCommand(utils.UQLCommand_createNodeProperty)
		break
	case types.DBType_DBEDGE:
		uql.SetCommand(utils.UQLCommand_createEdgeProperty)
		break
	}

	uql.SetCommandParams([]interface{}{propertyName, propertyType, description})

	res := t.UQLListSample(uql.ToString(), commonReq)

	return &ResponseCreateProperty{
		res.ResWithoutData,
	}
}

type ResponseDropProperty struct {
	*types.ResWithoutData
}

func (t *Connection) DropProperty(dataset ultipa.DBType, propertyName string, commonReq *types.Request_Common) *ResponseDropProperty {
	uql := utils.UQLMAKER{}
	switch dataset {
	case types.DBType_DBNODE:
		uql.SetCommand(utils.UQLCommand_dropNodeProperty)
		break
	case types.DBType_DBEDGE:
		uql.SetCommand(utils.UQLCommand_dropEdgeProperty)
		break
	}
	uql.SetCommandParams([]interface{}{propertyName})
	res := t.UQLListSample(uql.ToString(), commonReq)
	return &ResponseDropProperty{
		res.ResWithoutData,
	}
}

type RequestAlterProperty struct {
	PropertyName string
	Description  string
}
type ResponseAlterProperty struct {
	*types.ResWithoutData
}

func (t *Connection) AlterProperty(dataset ultipa.DBType, propertyName string, modified *RequestAlterProperty, commonReq *types.Request_Common) *ResponseAlterProperty {
	uql := utils.UQLMAKER{}
	switch dataset {
	case types.DBType_DBNODE:
		uql.SetCommand(utils.UQLCommand_alterNodeProperty)
		break
	case types.DBType_DBEDGE:
		uql.SetCommand(utils.UQLCommand_alterEdgeProperty)
		break
	}

	uql.SetCommandParams(propertyName)

	uql.AddParam("set", map[string]interface{}{
		"name":        modified.PropertyName,
		"description": modified.Description,
	}, true)

	log.Println(uql.ToString())
	res := t.UQLListSample(uql.ToString(), commonReq)

	return &ResponseAlterProperty{
		res.ResWithoutData,
	}
}
