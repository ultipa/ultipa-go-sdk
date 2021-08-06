package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/utils"
)

func TestListGraph(t *testing.T) {
	InitCases()
	res, err := client.ListGraph(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}

func TestCreateGraph(t *testing.T){

	client.DropGraph("zjs_amz", nil)

	resp, err := client.CreateGraph(&structs.Graph{
		Name: "zjs_amz",
		Description: "a graph",
	}, nil)

	assert.Equal(t, nil, err, "create graph request")
	assert.Equal(t, resp.Status.Code, ultipa.ErrorCode_SUCCESS, "create graph status")
}


func TestDeleteGraph(t *testing.T) {
	client.DropGraph("zjs_amz", nil)
}
