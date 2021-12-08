package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/printers"
)

func TestUQL(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.86:60072"}, "default")

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

		for _, a := range resp.AliasList {
			dataitem := resp.Alias(a)
			printers.PrintAny(dataitem)
			log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

		}
	}
}
