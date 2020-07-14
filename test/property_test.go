package test

import (
	"log"
	"testing"
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