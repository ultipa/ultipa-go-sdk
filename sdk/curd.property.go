package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ShowPropertyRequest = struct {
	Dataset types.DBType;
}

func (t *Connection) ListProperty (request ShowPropertyRequest) *types.ResAny {
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
	//urlData, ok := res.Data.(utils.UqlReply)
	_, ok := res.Data.(types.UqlReply)

	if ok {
		//properties := utils.TableToArray(urlData.Tables[0])
		//
		//for i := 0; i < len(*properties); i++ {
		//
		//}
		//var ps []*utils.Property
		//for _, pv := range *properties{
		//	pv := *pv
		//	p := utils.Property{
		//		PropertyName: pv["name"],
		//		PropertyType: pv["type"],
		//		Index: pv["index"] == "true",
		//		Lte: pv["lte"] == "true",
		//	}
		//	ps = append(ps, &p)
		//}
		//res.Data = ps
	}
	return &res
}