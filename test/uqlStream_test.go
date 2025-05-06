package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
)

func TestUQLStream(t *testing.T) {
	uql := `find().nodes() limit 10000 return nodes{*} `

	log.Println("Exec : ", uql)

	//resp, err := client.Uql(c.Uql, &configuration.RequestConfig{Graph: "multi_schema_test"})
	stream, err := client.UQLStream(uql, &configuration.RequestConfig{
		Graph: "alimama",
	})

	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for true {
		resp, err := stream.Recv(true)
		if err != nil {
			t.Fatal(err)
		}
		i++
		if resp != nil {
			printers.PrintStatistics(resp.Statistic)
			nodes, schema, err := resp.Get(0).AsNodes()
			if err != nil {
				t.Fatal(err)
			}
			printers.PrintNodes(nodes, schema)
			if i > 3 {
				stream.Recv(false)
				break
			}
		} else {
			break
		}
	}
	stream.Close()
}
