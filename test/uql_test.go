package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/printers"
)

func TestUQL(t *testing.T) {

	InitCases()

	for _, c := range cases {

		log.Println("Exec : ", c.UQL)

		//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
		resp, err := client.UQL(c.UQL, nil)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
			log.Println(resp.Status.Message)
			continue
		}

		for _, a := range c.Alias {
			dataitem := resp.Alias(a)
			printers.PrintAny(dataitem)

		}
	}
}
