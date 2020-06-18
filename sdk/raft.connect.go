package sdk

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type GetLeaderRequest struct {
}
type GetLeaderReply struct {
	Status *types.Status
}

func (t *Connection) GetLeaderReuqest() *GetLeaderReply {
	clientInfo, ctx, cancel := t.chooseClient(TIMEOUT_DEFAUL)
	defer cancel()
	res, err := clientInfo.Client.GetLeader(ctx, &ultipa.GetLeaderRequest{})
	//_json, _ := utils.StructToJSONString(res)
	//fmt.Printf(_json, err)
	resFormat := GetLeaderReply{
		Status: utils.FormatStatus(res.Status, err),
	}
	return &resFormat
}
