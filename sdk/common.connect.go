package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/types/types_response"
	"ultipa-go-sdk/utils"
)
func (t *Connection) Stat(commonReq *SdkRequest_Common) *types.ResStat {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.CommandList_stat)

	res := t.UQL(uql.ToString(),  commonReq)
	var newData types_response.Stat
	if res.Status.Code == types.ErrorCode_SUCCESS  {
		uqlReply := res.Data
		datas := utils.TableToArray((*uqlReply.Tables)[0])
		for _, data := range *datas{
			newData = types_response.Stat{
				MemUsage: (*data)["memUsage"].(string),
				CpuUsage: (*data)["cpuUsage"].(string),
			}
			break
		}
	}
	return &types.ResStat{
		res.ResWithoutData,
		&newData,
	}
}

func (t *Connection) ClusterInfo() *types.ResListClusterInfo {
	t.RefreshRaftLeader("",&SdkRequest_Common{
		GraphSetName: RAFT_GLOBAL,
	})
	hosts := t.HostManagerControl.GetAllHosts()
	var result []*types_response.ClusterInfo
	for _, host := range *hosts{
		host_commom_req := &SdkRequest_Common{
			UseHost: host,
		}
		status, _ := t.TestConnect(host_commom_req)
		info := &types_response.ClusterInfo{
			Status: status,
			Host: host,
			Stat: &types_response.Stat{
				MemUsage: "",
				CpuUsage: "",
			},
		}
		res := t.Stat(host_commom_req)
		if res.Status.Code == types.ErrorCode_SUCCESS {
			info.Stat = res.Data
		}
		result = append(result, info)
	}
	return &types.ResListClusterInfo{
		&types.ResWithoutData{
			Status: &types.Status{
				Code: types.ErrorCode_SUCCESS,
			},
		},
		result,
	}
}

