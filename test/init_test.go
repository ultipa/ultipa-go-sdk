package test

import (
	"testing"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
)

var client *api.UltipaAPI

func TestMain(m *testing.M) {



	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts: []string {
			"210.13.32.146:60074",
			//"localhost:8088",
		},
		Username: "root",
		Password: "root",
	})

	client = sdk.NewUltipa(config)
	m.Run()
}

