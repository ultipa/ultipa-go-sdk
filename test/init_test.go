package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
)

var client *api.UltipaAPI

func TestMain(m *testing.M) {
	var err error


	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts: []string {
			"210.13.32.146:60074",
			//"localhost:8088",
		},
		Username: "root",
		Password: "root",
		DefaultGraph: "multi_schema_test",
	})

	client, err = sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}




