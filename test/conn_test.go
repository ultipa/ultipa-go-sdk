package test

import (
	"testing"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/configuration"
)

func TestNewConn(t *testing.T) {
	config := configuration.UltipaConfig{
		Hosts: []string {
			"192.168.1.21:60074",
		},
		Username: "root",
		Password: "root",
	}

	sdk.NewUltipa(&config)
	//defer ultipa.Close()
}