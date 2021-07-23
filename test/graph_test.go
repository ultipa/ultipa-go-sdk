package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestListGraph(t *testing.T) {
	InitCases()
	res, err := client.ListGraph(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}