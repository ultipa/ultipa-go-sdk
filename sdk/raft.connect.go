package sdk
//
//import (
//	ultipa "ultipa-go-sdk/rpc"
//	"ultipa-go-sdk/types"
//	"ultipa-go-sdk/utils"
//)
//
//type GetLeaderRequest struct {
//}
//type GetLeaderReply struct {
//	Status *types.Status
//}
//
//func (t *Connection) GetLeaderReuqest(commonReq *types.Request_Common) *GetLeaderReply {
//	if commonReq == nil {
//		commonReq = &types.Request_Common{}
//	}
//	clientInfo := t.getClientInfo(&GetClientInfoParams{
//		ClientType: ClientType_Leader,
//		GraphSetName: commonReq.GraphSetName,
//		IgnoreRaft: true,
//	})
//	defer clientInfo.CancelFunc()
//	//fmt.Println("host", clientInfo.Host)
//	res, err := clientInfo.ClientInfo.Client.GetLeader(clientInfo.Context, &ultipa.GetLeaderRequest{})
//	//log.Printf(utils.StructToPrettyJSONString(res))
//	if err != nil {
//		return &GetLeaderReply{
//			Status: utils.FormatStatusWithHost(nil, err, clientInfo.Host),
//		}
//	} else {
//		return &GetLeaderReply{
//			Status: utils.FormatStatusWithHost(res.Status, nil, clientInfo.Host),
//		}
//	}
//}
