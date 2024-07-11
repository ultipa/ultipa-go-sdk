package test

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	res, err := client.ShowUser(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.DataItemMap)

	res, err = client.GetUser("yu", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.DataItemMap)
}
