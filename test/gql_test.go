package test

import "testing"

func TestGql(t *testing.T) {
	//response, err := client.Uql("find().nodes() as n return n{*}", nil)
	response, err := client.Gql("match (n) return n", nil)
	//response, err := client.Gql("CALL db.schema.visualization()", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(response.Status.Code.String())
}
