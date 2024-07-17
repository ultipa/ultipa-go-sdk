package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"log"
	"testing"
)

func TestDelete(t *testing.T) {
	conf := &configuration.InsertRequestConfig{
		RequestConfig: &configuration.RequestConfig{GraphName: "amz2"},
		Silent:        true,
	}
	resp, err := client.DeleteNodes("{_uuid < 200}", conf)
	if err != nil {
		log.Println(err, resp)
	}

	log.Println(resp.DataItemMap)

	resp, err = client.DeleteEdges("{_uuid < 200}", conf)
	if err != nil {
		log.Println(err, resp)
	}
	it := resp.Alias("edges")
	log.Println(it)
}
