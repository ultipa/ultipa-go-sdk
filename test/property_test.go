/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  property_test
 * @Date: 2022/7/29 5:45 下午
 */

package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
)

func TestShowProperty(t *testing.T) {
	resp, _ := client.UQL("show().property()", nil)

	nodeProperties, err := resp.Alias(http.RESP_NODE_PROPERTY_KEY).AsProperties()
	if err != nil {
		log.Fatalln(err)
	}
	edgeProperties, err := resp.Alias(http.RESP_EDGE_PROPERTY_KEY).AsProperties()
	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintProperty(nodeProperties)
	printers.PrintProperty(edgeProperties)
}

func TestShowNodeProperty(t *testing.T) {
	resp, _ := client.UQL("show().node_property()", nil)

	nodeProperties, err := resp.Alias(http.RESP_NODE_PROPERTY_KEY).AsProperties()
	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintProperty(nodeProperties)
}

func TestShowEdgeProperty(t *testing.T) {
	resp, _ := client.UQL("show().edge_property()", nil)

	edgeProperties, err := resp.Alias(http.RESP_EDGE_PROPERTY_KEY).AsProperties()
	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintProperty(edgeProperties)
}

func TestCreatePropertyWithUql(t *testing.T) {
	resp, err := client.UQL(`create().node_property(@People, "age", "int32[]")`, nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}

	resp, _ = client.UQL("show().node_property(@People)", nil)
	nodeProperties, err := resp.Alias(http.RESP_NODE_PROPERTY_KEY).AsProperties()
	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintProperty(nodeProperties)
}

func TestCreateProperty(t *testing.T) {
	// Create Node Property
	newProp := &structs.Property{
		Name: "gender",
		Type: ultipa.PropertyType_STRING,
	}

	resp, err := client.CreateProperty("People", ultipa.DBType_DBNODE, newProp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatalf("resp code:%v,message:%v", resp.Status.Code, resp.Status.Message)
	}
	log.Println(resp.Status.Code)
}
