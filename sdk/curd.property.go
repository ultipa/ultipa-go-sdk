package sdk

import (
	"ultipa-go-sdk/utils"
)

type ShowPropertyRequest = struct {
	Dataset DBType;
}

func (t *Connection) ListProperty (request ShowPropertyRequest) *Res{
	uql := utils.UQLMAKER{}
	dataset := request.Dataset
	switch dataset {
	case DBType_DBNODE:
		uql.SetCommand(utils.CommandList_showNodeProperty)
		break
	case DBType_DBEDGE:
		uql.SetCommand(utils.CommandList_showEdgeProperty)
		break
	}
	res := t.UQL(uql.ToString())
	urlData, ok := res.Data.(UqlReply)

	if ok {
		properties := TableToArray(urlData.Tables[0])

		for i := 0; i < len(*properties); i++ {

		}
		var ps []*Property
		for _, pv := range *properties{
			pv := *pv
			p := Property{
				PropertyName: pv["name"],
				PropertyType: pv["type"],
				Index: pv["index"] == "true",
				Lte: pv["lte"] == "true",
			}
			ps = append(ps, &p)
		}
		res.Data = ps
	}
	return &res
}