package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestListSchema(t *testing.T) {
	InitCases()
	res, err := client.ListNodeSchema(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}