package test

import (
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/utils"
)

func TestStringAsInterface(t *testing.T) {
	datetime, err := utils.StringAsInterface("1970-01-01", ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)
}
