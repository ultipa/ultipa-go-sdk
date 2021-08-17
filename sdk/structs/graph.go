package structs

import (
	"errors"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/types"
)

type Graph struct {
	ID          types.ID
	Name        string
	Description string
	TotalNodes  uint64
	TotalEdges  uint64
}

func GetDBTypeByString(str string) (ultipa.DBType, error) {

	switch strings.ToLower(str) {
	case "node":
		return ultipa.DBType_DBNODE, nil
	case "edge":
		return ultipa.DBType_DBEDGE, nil
	}

	return 0, errors.New("DBType is not Exist : " + str)
}

func DBTypeToString(dbType ultipa.DBType) string {

	switch dbType {
	case ultipa.DBType_DBNODE:
		return "node"
	case ultipa.DBType_DBEDGE:
		return "edge"
	}

	return ""
}
