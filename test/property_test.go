package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func TestGetProperty(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)

	res := conn.ListProperty(&sdk.ShowPropertyRequest{
		Dataset: types.DBType_DBNODE,
	}, nil)
	resJson, _ := utils.StructToJSONBytes(res)

	log.Printf("\nuql res ->\n %s\n", resJson)
	for _, pty := range res.Data {
		log.Printf("%s, %s, %v", pty.PropertyType, pty.PropertyName, pty.Lte)
	}
}

func TestCreateProperty(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)

	conn.CreateProperty(ultipa.DBType_DBNODE, "go_new_node_property", types.PROPERTY_TYPE_INT64_STRING, "desc of node 1", nil)
}

func TestDropProperty(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)

	conn.DropProperty(ultipa.DBType_DBNODE, "go_new_node_property", nil)
}

func TestAlterProperty(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)

	conn.CreateProperty(ultipa.DBType_DBNODE, "go_new_node_property", types.PROPERTY_TYPE_INT64_STRING, "desc of node 1", nil)

	conn.AlterProperty(ultipa.DBType_DBNODE, "go_new_node_property", &sdk.RequestAlterProperty{
		PropertyName: "go_new_node_property",
		Description:  "desc111",
	}, nil)
}
