package test

import (
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
)

func TestUQLStream(t *testing.T) {
	uql := `find().nodes({@nodx}) limit 1000000 return nodes{*} `

	log.Println("Exec : ", uql)

	//resp, err := client.Uql(c.Uql, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	stream, err := client.UQLStream(uql, nil)

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
