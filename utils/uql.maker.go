package utils

import (
	"fmt"
	"strings"
)

type UQLMAKER struct {
	_command  string
	_commandP string
	_params   []struct {
		key   string
		value string
	}
}

const (
	CommandList_ab                 string = "ab"
	CommandList_khop               string = "khop"
	CommandList_nodes              string = "find().nodes"
	CommandList_edges              string = "find().edges"
	CommandList_deleteNodes        string = "delete().nodes"
	CommandList_deleteEdges        string = "delete().edges"
	CommandList_updateNodes        string = "update().nodes"
	CommandList_updateEdges        string = "update().edges"
	CommandList_template           string = "t"
	CommandList_autoNet            string = "autoNet"
	CommandList_autoNetByPart      string = "autoNetByPart"
	CommandList_nodeSpread         string = "spread"
	CommandList_insertNode         string = "insert().nodes"
	CommandList_insertEdge         string = "insert().edges"
	CommandList_showProperty       string = "show().property"
	CommandList_showNodeProperty   string = "show().node_property"
	CommandList_showEdgeProperty   string = "show().edge_property"
	CommandList_createNodeProperty string = "create().node_property"
	CommandList_createEdgeProperty string = "create().edge_property"
	CommandList_dropNodeProperty   string = "drop().node_property"
	CommandList_dropEdgeProperty   string = "drop().edge_property"
	CommandList_lteNode            string = "LTE().node_property"
	CommandList_lteEdge            string = "LTE().edge_property"
	CommandList_ufeNode            string = "UFE().node_property"
	CommandList_ufeEdge            string = "UFE().edge_property"
	CommandList_createIndex        string = "createIndex"
	CommandList_showIndex          string = "showIndex"
	CommandList_dropIndex          string = "dropIndex"
	CommandList_stat               string = "stat"
	CommandList_algo               string = "algo"
	CommandList_listPrivilege      string = "listPrivilege"
	CommandList_grant              string = "grant"
	CommandList_revoke             string = "revoke"
	CommandList_listUser           string = "listUser"
	CommandList_getUser            string = "getUser"
	CommandList_createUser         string = "createUser"
	CommandList_updateUser         string = "updateUser"
	CommandList_deleteUser         string = "deleteUser"
	CommandList_createPolicy       string = "createPolicy"
	CommandList_updatePolicy       string = "updatePolicy"
	CommandList_deletePolicy       string = "deletePolicy"
	CommandList_listPolicy         string = "listPolicy"
	CommandList_getPolicy          string = "getPolicy"
	CommandList_showTask           string = "show().task"
	CommandList_clearTask          string = "clear().task"
)

func (t *UQLMAKER) SetCommand(command string) {
	t._command = command
}
func (t *UQLMAKER) SetCommandParams(commandP interface{}) {
	if nil == commandP {
		return
	}
	commandPString, ok := commandP.(string)
	if ok {
		t._commandP = commandPString
		return
	}
	jsonBytes, err := StructToJSONBytes(commandP)
	if nil != err {
		return
	}
	t._commandP = BytesToString(jsonBytes)
}

func (t *UQLMAKER) AddParam(key string, value interface{}, required bool)  {
	valueBool, ok := value.(bool)
	if ok {
		if valueBool {
			required = false
		}
		value = ""
	}
	if required {
		if nil == value {
			return
		}
 	}
	switch key {
	case "filter":
		t.AddParam("node_filter", value, true);
		t.AddParam("edge_filter", value, true);
		return;
	case "select":
	// t.AddParam("select_node_properties", value, true);
	// t.AAddParam("select_edge_properties", value, true);
	// return;
	case "select_node_properties":
	case "select_edge_properties":
	case "privileges":
	case "policies":
		valueArray, ok := value.([]string)
		if ok {
			value = strings.Join(valueArray, ",")
		}
		break;
	}
	valueString, ok := value.(string)
	if !ok {
		bytes, err := StructToJSONBytes(value)
		if nil != err {
			return
		}
		valueString = BytesToString(bytes)
	}
	t._params = append(t._params, struct {
		key   string
		value string
	}{key: key, value: valueString})
}
func (t *UQLMAKER) ToString() string {
	uql := ""
	uql = fmt.Sprintf("%s%s(%s)", uql,t._command,t._commandP)
	if len(t._params) > 0 {
		var strs []string
		for _, v := range t._params {
			strs = append(strs, fmt.Sprintf("%s(%s)", v.key, v.value))
		}
		uql = uql + strings.Join(strs, ".")
	}
	return uql
}



