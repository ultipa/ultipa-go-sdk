package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

//1.86 1.85 1.90

func TestRefreshClusterInfo(t *testing.T) {

	hosts := []string{
		"192.168.1.86:40101",
		"192.168.1.85:40101",
		"192.168.1.90:40101",
	}
	client, err := GetClient(hosts, "default")

	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 10;  i++ {
		client.Pool.RefreshClusterInfo("global")
		utils.PrintJSON(client.Pool.GraphInfos)
	}

}

func TestSendNewGraphUQL(t *testing.T) {

}
