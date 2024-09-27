package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"io"
	"testing"
)

func TestGql(t *testing.T) {
	//response, err := client.Uql("find().nodes() as n return n{*}", nil)
	response, err := client.Gql("match (n) return n", nil)
	//response, err := client.Gql("CALL db.schema.visualization()", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(response.Status.Code.String())
}

func TestGqlStream(t *testing.T) {
	stream, err := client.GQLStream("match (n) return n", nil)

	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for true {
		resp, err := stream.Recv(true)
		if err == io.EOF {
			break
		} else if err != nil {
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
