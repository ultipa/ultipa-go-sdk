package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"log"
	"testing"
)

func TestUQLStream(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	client, _ := GetClient(hosts, graph)

	uql := `find().nodes({@nodx}) limit 1000000 return nodes{*} `

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	stream, err := client.UQLStream(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	i := 0
	for true {
		resp, err := stream.Recv(true)
		if err != nil {
			log.Fatalln(err)
		}
		i++
		if resp != nil {
			printers.PrintStatistics(resp.Statistic)
			nodes, schema, err := resp.Get(0).AsNodes()
			if err != nil {
				log.Fatalln(err)
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
