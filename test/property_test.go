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
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/printers"
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
