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

func (t *Connection) GetLeaderReuqest(req *SdkRequest_Common) *GetLeaderReply {
	if req == nil {
		req = &SdkRequest_Common{}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		ClientType: ClientType_Leader,
		GraphSetName: req.GraphSetName,
		IgnoreRaft: true,
	})
	defer clientInfo.CancelFunc()
	//fmt.Println("host", clientInfo.Host)
	res, err := clientInfo.ClientInfo.Client.GetLeader(clientInfo.Context, &ultipa.GetLeaderRequest{})
	//log.Printf(utils.StructToPrettyJSONString(res))
	if err != nil {
		return &GetLeaderReply{
			Status: utils.FormatStatus(nil, err),
		}
	} else {
		return &GetLeaderReply{
			Status: utils.FormatStatus(res.Status, nil),
		}
	}
}
